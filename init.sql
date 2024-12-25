drop database if exists bi_activity;
create database bi_activity;
use bi_activity;

ALTER TABLE student MODIFY COLUMN student_id VARCHAR(30);

select college.campus from college where campus = '' or college_account = ''
