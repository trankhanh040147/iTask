/* dbdiagram.io Database Schema */

Table Users {
    id bigserial [pk] 
    username varchar [not null]
    email varchar [not null]
    full_name varchar [not null]
    role bigint [not null, ref: > Roles.code]
    status int [not null]
    password_hash varchar [not null]
    salt varchar [not null]
    address varchar
    phone varchar
    dob varchar
    profile_ava_url varchar
    profile_cover_url varchar
    is_email_verified int
    bio varchar
    created_at timestampz [default: `now()`]
    updated_at timestampz [default: `now()`]
}

Table Projects {
    id bigserial [pk] 
    name varchar [not null]
    description varchar
    status int [not null]
    thumbnail_url varchar   
    priority int 
    privacy int 
    created_by bigint [not null, ref: > Users.id]
    deadline timestamp  
    started_at timestamp
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
    thumbnail varchar
}

Table Tasks {
    id bigserial [pk] 
    parent_task_id int 
    project_id bigint [not null, ref: > Projects.id]
    status int [not null]
    created_by bigint [not null, ref: > Users.id]
    name varchar [not null]
    description varchar
    position float
    priority int
    completed bool [not null]
    due_date timestamp
    started_at timestamp 
    completed_at timestamp
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

Table Attachments {
    id bigserial [pk] 
    task_id int [ref: > Tasks.id]
    project_id bigint [ref: > Projects.id]
    created_by bigint [ref: > Users.id]
    file_name varchar
    file_path varchar
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

Table Tags {
    id bigserial [pk] 
    tag_type int /* 1: task, 2: project */
    name varchar    
    description varchar
    position float
}

Table ProjectTags {
    id bigserial [pk] 
    project_id bigint [ref: > Projects.id]
    tag_id bigint [ref: > Tags.id]
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

Table TaskTags {
    id bigserial [pk] 
    task_id bigint [ref: > Tasks.id]
    tag_id bigint [ref: > Tags.id]
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

Table Notifications {
    id bigserial [pk] 
    caused_by bigint [ref: > Users.id]
    action_type varchar
    data json
    created_on timestamp [default: `now()`]
}

Table NotificationNotified {
    id bigserial [pk] 
    notification_id bigint [ref: > Notifications.id]
    user_id bigint [ref: > Users.id]
    read bool
    read_at timestamp [default: `now()`]
}

Table ProjectMembers {
    id bigserial [pk] 
    project_id bigint [ref: > Projects.id]
    user_id bigint [ref: > Users.id]
    added_at timestamp [default: `now()`]
    role_code varchar
}

Table ProjectMemberInvited {
    id bigserial [pk] 
    project_id bigint [ref: > Projects.id]
    user_account_invited_id bigint [ref: > Users.id]
}

Table Roles {
    code varchar [pk]
    name varchar
}

/*
Table Activities {
    id bigserial [pk] 
    project_id bigint [ref: > Projects.id]
    task_id bigint [ref: > Tasks.id]
    user_id bigint [ref: > Users.id]
    content varchar
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}
*/

Table TaskActivities {
    id bigserial [pk] 
    active bool
    task_id bigint [ref: > Tasks.id]
    created_at timestamp [default: `now()`]
    caused_by bigint [ref: > Users.id] 
    activity_type_id bigint [ref: > TaskActivityTypes.id]
    data json
}

Table TaskActivityTypes {
    id bigserial [pk]
    code varchar
    template varchar
}

Table TaskAssigned {
    id bigserial [pk] 
    task_id bigint [ref: > Tasks.id]
    user_id bigint [ref: > Users.id]
    assigned_date timestamp [default: `now()`]
}

Table TaskWatchers {
    id bigserial [pk] 
    task_id bigint [ref: > Tasks.id]
    user_id bigint [ref: > Users.id]
    watched_at timestamp [default: `now()`]
}
