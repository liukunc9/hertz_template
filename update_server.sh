#!/bin/bash

for file in ./http_idl/*; do
  if [ -f "$file" ]; then
    hz update -handler_dir ./handler -model_dir ./model -router_dir ./router -module github.com/liukunc9/hertz_template/ -idl $file
  fi
done