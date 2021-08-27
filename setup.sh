#!/bin/sh

if ! command -v ruby &> /dev/null 
then
  echo "Cannot find ruby, please install ruby"
  exit -127
fi

if ! command -v gem &> /dev/null 
then 
  echo "Please install rubygems"
  exit -127
fi

gem install bundler

bundle install

yarn install
