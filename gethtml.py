from requests_html import HTMLSession 
import h2md
def get_html(): 
    session = HTMLSession()
    r = session.get('http://www.imooc.com/wiki/javalesson/operators.html')
    content = r.html.find('.content',first = True)
    file = 'operators.md'
    
    text = h2md.convert(content.html)
    file_data = ""
    # 写入处理后的内容
    with open(file, 'w') as f:
        f.write(text)
    
    for line in open(file): 
        if line.find("运行案例") == -1 and line.find("实例演示") == -1 and line.find("复制") == -1:
            file_data += line

    with open(file,"w",encoding="utf-8") as f:
        f.write(file_data)


if __name__ == '__main__':
    get_html()