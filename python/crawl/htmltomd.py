import html2text as ht
import os
import requests


def tomd():
    
    path = os.getcwd()
    # 读取html格式文件
    with open(path+'/未更新完/解读你身边的经济学/7794.html', 'r', encoding='UTF-8') as f:
        htmlpage = f.read()
    # 处理html格式文件中的内容
    text = ht.HTML2Text().handle(htmlpage)
    # 写入处理后的内容
    with open('7794.md', 'w') as f:
        f.write(text)
if __name__ == '__main__':
    tomd()