#!/bin/sh

HOST='reef.moris.tech'
USER='reef'
DOC_ROOT="/home/reef"




if [ $USER ]
then
  SSH_HOST="$USER@$HOST"
else
  SSH_HOST=$HOST
fi

echo '* Создаем архив...'
tar -czf yourproject.tar.gz reef version.json templates/* static/*
if [ $? -ne 0 ]
then
  exit 1;
fi;



echo '* Копируем архив на сервер...'

scp -P '42401' ./yourproject.tar.gz  $SSH_HOST:$DOC_ROOT
if [ $? -ne 0 ]
then
  exit 1;
fi;





echo '* Распаковываем архив на серверe...'
#ssh -p 42401 $SSH_HOST "cd $DOC_ROOT; tar -xzf yourproject.tar.gz 2> /dev/null && rm -rf $DOC_ROOT/goapp/reef && rm -rf $DOC_ROOT/goapp/version.json  && rm -rf $DOC_ROOT/goapp/templates/* && rm -rf $DOC_ROOT/public_html/static/* && rm -rf $DOC_ROOT/goapp/static/* &&  mv templates/ $DOC_ROOT/goapp/ && mv static/ $DOC_ROOT/public_html/ &&  mv reef $DOC_ROOT/goapp  && mv version.json $DOC_ROOT/goapp && chmod -R a+w+x $DOC_ROOT/goapp/reef"
ssh -p 42401 $SSH_HOST "cd $DOC_ROOT; tar -xzf yourproject.tar.gz 2> /dev/null && rm -rf $DOC_ROOT/goapp/reef && rm -rf $DOC_ROOT/goapp/version.json  &&  mv reef $DOC_ROOT/goapp  && mv version.json $DOC_ROOT/goapp && chmod -R a+w+x $DOC_ROOT/goapp/reef"
if [ $? -ne 0 ]
then
  exit 1;
fi;



echo '* Удаляем архив на сервере ...'
ssh -p 42401 $SSH_HOST "cd $DOC_ROOT; rm -rf yourproject.tar.gz"
if [ $? -ne 0 ]
then
  exit 1;
fi;

echo '* Удаляем архив локально ...'

rm -rf yourproject.tar.gz