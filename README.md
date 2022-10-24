# go-wiki-homework

URL=https://go.dev/doc/articles/wiki/  
#### 前書き  
HTML,CSS,GOやGitHub自体まったくと言っていいほど触ったことがなく  
不手際や不適切な部分がある可能性が高いため前もってお詫び致します．  
お忙しいところ恐縮ですがお見通し頂けたら幸いです．  

## task1

#### 課題内容  
Other tasks  
Here are some simple tasks you might want to tackle on your own:  

1.1 Store templates in tmpl/ and page data in data/.  
1.2 Add a handler to make the web root redirect to /view/FrontPage.  
1.3 Spruce up the page templates by making them valid HTML and adding some CSS rules.  
1.4 Implement inter-page linking by converting instances of [PageName] to
<a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)  

#### 実行方法  
cmd 
>cd %PATH%  
>go build assignment.go  
>assignment.exe  

## task2
#### 課題内容  
task1 のassignment.goにて以下のURLから main_testをコピーしテストを実行する  
https://gist.github.com/ymt2/06ae6f7f9a35224eb477e1ca72fb0f52  
#### 実行方法  
cmd  
>cd %PATH%  
>go test main_test.go assignment.go  

## task3
#### 課題内容  
task2 のassignment.go,をmain_testにて-raceを用いて競合のテストを行う  
#### 実行方法  
cmd  
>cd %PATH%  
>go test -race main_test.go assignment.go  
