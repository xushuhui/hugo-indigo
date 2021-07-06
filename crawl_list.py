#encoding:utf-8
import requests
from const import headers 

# 获取所有专栏并且一键订阅所有未购买专栏
def lessions_subscription():
    lession_url = 'https://gate.lagou.com/v1/neirong/edu/homepage/getCourseListV2?isPc=true'
    pay_url = 'https://gate.lagou.com/v1/neirong/edu/member/drawCourse?courseId={id}'
    
   
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
