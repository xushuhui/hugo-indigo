
#encoding:utf-8
import os
def merge():
    path = "./未更新完/"
    rpath= "./merge/"
    filenames = os.listdir(path)
   
    for filename in filenames:
        fn = os.listdir(path + filename)
        
        result = rpath + filename+".html"
        file = open(result, 'w+', encoding="utf-8")
        for f in fn:
            
            for line in open(path + filename+"/"+f, encoding="utf-8"):
                if line.find("金句") != -1:
                    continue
                file.writelines(line)
            file.write('\n')
if __name__ == '__main__':
    merge()