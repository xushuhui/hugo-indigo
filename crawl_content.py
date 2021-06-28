#encoding:utf-8
import asyncio
import logging
import os
from multiprocessing import Pool

import aiohttp
import requests
from aiohttp import ContentTypeError
from jinja2 import Environment, FileSystemLoader

logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s - %(levelname)s: %(message)s')
LESSION_URL = 'https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons?courseId={id}'
INDEX_URL = 'https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail?lessonId={id}'
COURSE_LIST_URL = 'https://gate.lagou.com/v1/neirong/edu/homepage/getCourseListV2?isPc=true'
PAY_URL = 'https://gate.lagou.com/v1/neirong/edu/member/drawCourse?courseId={id}'
PURCHASE_URL = 'https://gate.lagou.com/v1/neirong/kaiwu/getAllCoursePurchasedRecordForPC'
CONCURRENCY = 5

loop = asyncio.get_event_loop()

headers = {
    'Host': 'gate.lagou.com',
    'Connection': 'keep-alive',
    'Pragma': 'no-cache',
    'Cache-Control': 'no-cache',
    'sec-ch-ua': '"Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"',
    'Accept': 'application/json, text/plain, */*',
    'Authorization': '50dc5543e28179e1b1f19b2b33dd822d4a77358f1ad63131',
    'X-L-REQ-HEADER': '{"deviceType":1,"userToken":"50dc5543e28179e1b1f19b2b33dd822d4a77358f1ad63131"}',
    'sec-ch-ua-mobile': '?0',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36',
    'Origin': 'https://edu.lagou.com',
    'Sec-Fetch-Site': 'same-site',
    'Sec-Fetch-Mode': 'cors',
    'Sec-Fetch-Dest': 'empty',
    'Referer': 'https://edu.lagou.com/',
    'Accept-Encoding': 'gzip, deflate, br',
    'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8',
    'Cookie': ''
   }


def lessions_list():
    response = requests.get(url=COURSE_LIST_URL, headers=headers)
    lessions_all = []
    if response is not None:
        course_list = response.json()['content']['contentCardList']
        for courses in course_list:
            if courses['cardType'] == 201:
                lessions_all = courses['courseList']
    return lessions_all


class Spider(object):

    def __init__(self):
        self.session = aiohttp.ClientSession()
        self.semaphore = asyncio.Semaphore(CONCURRENCY)

    def scrape_lession_content(self, url, course_id):
        url = url.format(id=course_id)
        logging.info('scraping %s', url)
        param = {'courseId': course_id}
        r = requests.get(url, params=param, headers=headers)
        return r.json()

    async def scrape_api(self, url):
        async with self.semaphore:
            try:
               
                params = {'lessonId': url.split('=')[1]}
                async with self.session.get(url, headers=headers, params=params) as response:
                    await asyncio.sleep(1)
                    return await response.json()
            except ContentTypeError as e:
                logging.error('error occurred while scraping %s', url, exc_info=True)

    async def scrape_detail(self, id):
        url = INDEX_URL.format(id=id)
        return await self.scrape_api(url)

    # html名称,目录名称,pdf内容,pdf所在文件夹
    async def write_html(self, file_name, theme, data, file_path):
        
        env = Environment(loader=FileSystemLoader('./'))
        template = env.get_template('template.htm')

        if data:
            xlsx_path = os.getcwd() + file_path
            if not os.path.exists(xlsx_path):
                os.makedirs(xlsx_path)
            with open(xlsx_path + '/' + file_name, 'w+', encoding='utf-8') as fout:
                html_content = template.render(theme=theme, body=data)
                fout.write(html_content)

    async def main(self, courceId):
        response = self.scrape_lession_content(LESSION_URL, course_id=courceId)
        course_section_list = response['content']['courseSectionList']
        lession_id = []
        lession_status = True
        file_path = response['content']['courseName']
        for course in course_section_list:
            for lession in course['courseLessons']:
                lession_status = lession['status']
                if lession_status == "RELEASE":
                    lession_id.append(lession['id'])
                elif lession_status == "UNRELEASE":
                    lession_status = False
        scrape_index_tasks = [asyncio.ensure_future(self.scrape_detail(page)) for page in lession_id]
        results = await asyncio.gather(*scrape_index_tasks)
        
        contents = []
        for index_data in results:
            content = index_data['content']
            if content is None:
                continue
            data = {'id': int(content['id']), 'theme': str(content['theme']),
                    'content': str(content['textContent'])}
            contents.append(data)

        output_path = ''
        if lession_status:
            output_path = '/' + '已更新完/' + file_path
        else:
            f = 'unreleased.txt'
            with open(f, "a") as file:
                file.write(str(courceId) + "\n")
            output_path = '/' + '未更新完/' + file_path + '（未更新完）'

        scrape_detail_tasks = [asyncio.ensure_future(self.write_html(str(content['id']) + '.html',
                                                                     content['theme'], content['content'],
                                                                     output_path))
                               for content in contents]
        await asyncio.wait(scrape_detail_tasks)
        await self.session.close()

    @classmethod
    def crawl_all(cls):
        # 全量爬取
        lessions = lessions_list()
        pool = Pool(processes=5)
        for lession in lessions:
            
            if lession['tag'] != '上新优惠':
                lession_id = lession['id']
                spider = Spider()
                pool.apply_async(loop.run_until_complete(spider.main(lession_id)))
                print('======> 开始爬取专栏：{}，编号：{} <======'.format(lession['title'], lession_id))
        pool.join()
        pool.close()

    @classmethod
    def crawl_increase(cls):
        # 增量爬取
        unreleases = set([line.rstrip('\n') for line in open('unreleased.txt', 'r')])
        lessions = list(unreleases | set(
            [line.rstrip('\n') for line in open('downloads.txt', 'r')]))
        pool = Pool(processes=5)
        open("unreleased.txt", 'w').close()
        for lession_id in lessions:
            spider = Spider()
            print('======> 开始爬取编号：{} <======'.format(lession_id))
            pool.apply_async(loop.run_until_complete(spider.main(lession_id)))
        pool.close()
        pool.join()
        open("downloads.txt", 'w').close()
        updated_unreleased_course = set([line.rstrip('\n') for line in open('unreleased.txt', 'r')])
        print('部分专栏由未完成变为已完成{}'.format(list(unreleases - updated_unreleased_course)))
        unreleased_courses = list(unreleases & updated_unreleased_course)
        open("unreleased.txt", 'w').close()
        with open('unreleased.txt', "a") as file:
            for unreleased_course in unreleased_courses:
                file.write(str(unreleased_course) + "\n")

def purchase_lessions():
    response = requests.get(url=PURCHASE_URL, headers=headers)
    lessions_all = []
    if response is not None:
       
        course_list = response.json()['content']['allCoursePurchasedRecord']
        for courses in course_list:
            lessions_all = courses['courseRecordList']
    return lessions_all
if __name__ == '__main__':
    
    spider = Spider()
    # 全量爬取
    spider.crawl_all()
    # 增量爬取
    #spider.crawl_increase()
