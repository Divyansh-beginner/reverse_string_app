#!/bin/bash
echo "-------------- $(date) -----------------"
echo "daily server setup at 7pm"
cd /home/divya_singh/assignments_folder/assignment_3_go/reverse_string_app
echo "moved in the required directory"
git checkout addingservertodatabase
echo "moved to database branch"
go run . &
echo "ran the files go run . & used"
echo "--------------------------------------------"



