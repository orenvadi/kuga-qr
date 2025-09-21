here is my postgres db diagram written in mermaid


```mermaid
erDiagram
  admin{
    int id PK
    string password_hash
  }

  the_user{
    string id PK
    string full_name
    string password_hash
  }

  teacher{
    string id FK
  }
  teacher||--||the_user : foregn-key


  student{
    string id FK
    int semester
  }
  student||--||the_user : foregn-key

  the_group{
    string id PK
    string student_id FK
  }
  the_group||--|{student: contains-many

  subject{
    string id PK
    string teacher_id FK
  }
  subject||--|{teacher: contains-many

  room{
    string id pk
    string room_code
  }

  schedule{
    int id PK
    string subject_id FK
    string group_id FK
    string teacher_id FK
    string room_id FK
    timestamp time
  }
  schedule||--||teacher: contains-one
  schedule||--||the_group: contains-one
  schedule||--||teacher: contains-one
  schedule||--||room: contains-one

  course{
    int id PK
    string group_id FK
    string subject_id FK
    string teacher_id FK
  }
  course||--||the_group: contains-one
  course||--||subject: contains-one
  course||--||teacher: contains-one

  marks_and_absence{
    int schedule_id FK
    string student_id FK
    int mark
    string status
  }
  marks_and_absence||--||schedule: contains-one
  marks_and_absence||--|{student: contains-many

  colloqium{
    int id PK
    int year
    int semester
    int course_id FK
    int colloq_1_score
    int colloq_2_score
    int colloq_3_score
    int exam_score
    int additional_score
    int total_score
  }
  colloqium||--||course: contains-one

  course_work{
    int id PK
    int year
    int semester
    int course_id FK
    int score
  }
  course_work||--||course: contains-one
```


here is the essential handlers i want

```
# Handlers

## Teacher
- login
- get_teacher_schedule
- stream_qr_codes

## Student
- login
- get_student_schedule
- scan_and_send_qr_codes
```



based on this data

write openapi.yaml file to generate server in go language and server in ts
