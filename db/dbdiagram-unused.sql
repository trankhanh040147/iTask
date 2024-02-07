/* dbdiagram.io Database Schema */

/* ! Alternatively for Activities feature
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

Table user_account_confirm_token {
    confirm_token_id uuid [pk, default: `uuid_generate_v4()`]
    email text [not null, unique]
}

Table TaskWatchers {
    id bigserial [pk] 
    task_id bigint [ref: > Tasks.id]
    user_id bigint [ref: > Users.id]
    watched_at timestamp [default: `now()`]
}
