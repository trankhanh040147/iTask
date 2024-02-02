/* dbdiagram.io Database Schema */

Table Users {
    id bigserial [pk] /* auto-increment */
    username varchar [not null]
    email varchar [not null]
    full_name varchar [not null]
    role int [not null, ref: > Roles.code]
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
    id bigserial [pk] /* auto-increment */
    name varchar [not null]
    description varchar
    status int [not null]
    priority int 
    privacy int 
    created_by int [not null, ref: > Users.id]
    deadline timestamp  
    started_at timestamp
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
    thumbnail varchar
}

Table Tasks {
    id bigserial [pk] /* auto-increment */
    parent_task_id int 
    project_id int [not null, ref: > Projects.id]
    status int [not null]
    created_by int [not null, ref: > Users.id]
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
    id bigserial [pk] /* auto-increment */
    task_id int [ref: > Tasks.id]
    project_id int [ref: > Projects.id]
    user_id int [ref: > Users.id]
    file_name varchar
    file_path varchar
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

Table Tags {
    id bigserial [pk] /* auto-increment */
    tag_type int /* 1: task, 2: project */
    name varchar    
    description varchar
    position float
}

Table ProjectTags {
    id bigserial [pk] /* auto-increment */
    project_id int [ref: > Projects.id]
    tag_id int [ref: > Tags.id]
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

Table TaskTags {
    id bigserial [pk] /* auto-increment */
    task_id int [ref: > Tasks.id]
    tag_id int [ref: > Tags.id]
    created_at timestamp [default: `now()`]
    updated_at timestamp [default: `now()`]
}

-- Table Activities {
--     id bigserial [pk] /* auto-increment */
--     project_id int [ref: > Projects.id]
--     task_id int [ref: > Tasks.id]
--     user_id int [ref: > Users.id]
--     content varchar
--     created_at timestamp [default: `now()`]
--     updated_at timestamp [default: `now()`]
-- }

Table Notifications {
    id bigserial [pk] /* auto-increment */
    caused_by int [ref: > Users.id]
    action_type varchar
    data json
    created_on timestamp [default: `now()`]
}

Table NotificationNotified {
    id bigserial [pk] /* auto-increment */
    notification_id int [ref: > Notifications.id]
    user_id int [ref: > Users.id]
    read bool
    read_at timestamp [default: `now()`]
}

Table ProjectMembers {
    id bigserial [pk] /* auto-increment */
    project_id int [ref: > Projects.id]
    user_id int [ref: > Users.id]
    added_at timestamp [default: `now()`]
    role_code varchar
}

Table ProjectMemberInvited {
    id bigserial [pk] /* auto-increment */
    project_id int [ref: > Projects.id]
    user_account_invited_id int [ref: > Users.id]
}

Table Roles {
    code varchar [pk]
    name varchar
}

Table TaskActivities {
    id bigserial [pk] /* auto-increment */
    active bool
    task_id int [ref: > Tasks.id]
    created_at timestamp [default: `now()`]
    caused_by int [ref: > Users.id]
    activity_type_id int
    data json
}

Table TaskActivityTypes {
    id bigserial [pk] /* auto-increment */
    code varchar
    template varchar
}

Table TaskAssigned {
    id bigserial [pk] /* auto-increment */
    task_id int [ref: > Tasks.id]
    user_id int [ref: > Users.id]
    assigned_date timestamp [default: `now()`]
}

Table TaskWatchers {
    id bigserial [pk] /* auto-increment */
    task_id int [ref: > Tasks.id]
    user_id int [ref: > Users.id]
    watched_at timestamp [default: `now()`]
}
