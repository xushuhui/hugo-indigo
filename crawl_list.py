#encoding:utf-8
import requests


# 获取所有专栏并且一键订阅所有未购买专栏
def lessions_subscription():
    lession_url = 'https://gate.lagou.com/v1/neirong/edu/homepage/getCourseListV2?isPc=true'
    pay_url = 'https://gate.lagou.com/v1/neirong/edu/member/drawCourse?courseId={id}'
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
   
    response = requests.get(url=lession_url, headers=headers)
    lessions_all = []
    if response is not None:
        # print(response.json())
        lessions_all = []
        course_list = response.json()['content']['contentCardList']
        for courses in course_list:
            if courses['cardType'] == 201:
                lessions_all = courses['courseList']
    lession_size = len(lessions_all)
    
    for lession in lessions_all:
        if not lession['hasBuy']:
            pay_response = requests.get(url=pay_url.format(id=lession['id']), headers=headers)
           
            if lession['tag'] != '上新优惠':
                if pay_response.json()['content'] is None:
                    lession_size = lession_size - 1
                    print('课程', lession['title'], 'vip 暂时无法订购1！！！')
                else:
                    with open('downloads.txt', "a") as file:
                        file.write(str(lession['id']) + "\n")
                    print('课程', lession['title'], str(lession['id']), '订购成功！！！')
            else:
                lession_size = lession_size - 1
                print('课程', lession['title'], 'vip 暂时无法订购2！！！')
        else:
            print('课程', lession['title'], '已订阅！！！')
    print('恭喜，拉钩专栏当前拥有{}，已经全部订阅！！！'.format(lession_size))


if __name__ == '__main__':
    lessions_subscription()
