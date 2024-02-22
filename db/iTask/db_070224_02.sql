CREATE TABLE `Users`
(
    `id`                BIGINT PRIMARY KEY,
    `username`          varchar(255) NOT NULL,
    `email`             varchar(255) NOT NULL,
    `full_name`         varchar(255) NOT NULL,
    `role_code`         int          NOT NULL,
    `title`             varchar(255),
    `status`            int          NOT NULL,
    -- 1: active, 0: banned/deleted
    `password_hash`     varchar(255) NOT NULL,
    `salt`              varchar(255) NOT NULL,
    `address`           varchar(255),
    `phone`             varchar(255),
    `dob`               date,
    `profile_ava_url`   varchar(255),
    `profile_cover_url` varchar(255),
    `is_email_verified` int,
    `bio`               varchar(255),
    `created_at`        timestamp DEFAULT (now()),
    `updated_at`        timestamp DEFAULT (now())
);

CREATE TABLE `Projects`
(
    `id`            BIGINT PRIMARY KEY,
    `name`          varchar(255) NOT NULL,
    `description`   varchar(255),
    `status`        int          NOT NULL,
    -- 1: active, 2: archived
    `thumbnail_url` varchar(255),
    `priority`      int,
    -- 1: low, 2: medium, 3: high
    `privacy`       int,
    -- 1: public, 2: private
    `created_by`    bigint       NOT NULL,
    `deadline`      timestamp,
    `started_at`    timestamp,
    `created_at`    timestamp DEFAULT (now()),
    `updated_at`    timestamp DEFAULT (now()),
    `thumbnail`     varchar(255)
);

CREATE TABLE `Tasks`
(
    `id`             BIGINT PRIMARY KEY,
    `parent_task_id` int,
    `project_id`     bigint       NOT NULL,
    `status`         int          NOT NULL,
    -- 1: to do, 2: in progress, 3: done, 4: deleted
    `created_by`     bigint       NOT NULL,
    `name`           varchar(255) NOT NULL,
    `description`    varchar(255),
    `position`       float,
    `priority`       int,
    `completed`      bool      DEFAULT false,
    `due_date`       timestamp,
    `started_at`     timestamp,
    `completed_at`   timestamp,
    `created_at`     timestamp DEFAULT (now()),
    `updated_at`     timestamp DEFAULT (now())
);

CREATE TABLE `TaskAttachments`
(
    `id`         BIGINT PRIMARY KEY,
    `task_id`    int,
    `created_by` bigint,
    `downloads`  int,
    `file_name`  varchar(255),
    `file_path`  varchar(255),
    `created_at` timestamp DEFAULT (now()),
    `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `Tags`
(
    `id`          BIGINT PRIMARY KEY,
    `tag_type`    int,
    -- 1: project, 2: task
    `name`        varchar(255),
    `description` varchar(255),
    `position`    float,
    `created_at`  timestamp DEFAULT (now()),
    `updated_at`  timestamp DEFAULT (now())
);

CREATE TABLE `ProjectTags`
(
    `id`         BIGINT PRIMARY KEY,
    `project_id` bigint,
    `tag_id`     bigint,
    `created_at` timestamp DEFAULT (now()),
    `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `TaskTags`
(
    `id`         BIGINT PRIMARY KEY,
    `task_id`    bigint,
    `tag_id`     bigint,
    `created_at` timestamp DEFAULT (now()),
    `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `Notifications`
(
    `id`          BIGINT PRIMARY KEY,
    `caused_by`   bigint,
    `action_type` varchar(255),
    `data`        json,
    `created_on`  timestamp DEFAULT (now())
);

CREATE TABLE `NotificationNotified`
(
    `id`              BIGINT PRIMARY KEY,
    `notification_id` bigint,
    `user_id`         bigint,
    `read`            bool,
    `read_at`         timestamp DEFAULT (now())
);

CREATE TABLE `ProjectMembers`
(
    `project_id` bigint,
    `user_id`    bigint,
    `added_at`   timestamp DEFAULT (now()),
    `role`       int,
    PRIMARY KEY (`project_id`, `user_id`)
);

CREATE TABLE `ProjectMemberInvited`
(
    `id`                      BIGINT PRIMARY KEY,
    `project_id`              bigint,
    `user_account_invited_id` bigint
);

CREATE TABLE `UserAccountInvited`
(
    `id`         BIGINT PRIMARY KEY,
    `email`      text      NOT NULL,
    `invited_on` timestamp NOT NULL DEFAULT (now()),
    `has_joined` boolean   NOT NULL DEFAULT false
);

CREATE TABLE `Roles`
(
    `code` int PRIMARY KEY,
    `name` varchar(255)
);

CREATE TABLE `TaskAssigned`
(
    `id`            BIGINT PRIMARY KEY,
    `task_id`       bigint,
    `user_id`       bigint,
    `assigned_date` timestamp DEFAULT (now())
);

CREATE TABLE `AuthTokens`
(
    `id`         BIGINT PRIMARY KEY,
    `token_id`   varchar(255) NOT NULL,
    `user_id`    bigint       NOT NULL,
    `created_at` timestamp DEFAULT (now()),
    `expires_at` timestamp    NOT NULL
);

CREATE TABLE `TaskComments`
(
    `id`                BIGINT PRIMARY KEY,
    `task_id`           bigint,
    `parent_comment_id` bigint,
    `created_by`        bigint,
    `message`           varchar(255),
    `pinned`            bool,
    `created_at`        timestamp DEFAULT (now()),
    `updated_at`        timestamp DEFAULT (now())
);

CREATE TABLE verify_emails
(
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    email      varchar(255) NOT NULL,
    scret_code varchar(255) NOT NULL,
    type       int          NOT NULL,
    created_at timestamp DEFAULT (now()),
    expired_at timestamp    NOT NULL
)

-- auto increment
ALTER TABLE
    `Users`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `Projects`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `Tasks`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `TaskAttachments`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `Tags`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `ProjectTags`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `TaskTags`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `Notifications`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `NotificationNotified`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `ProjectMembers`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `ProjectMemberInvited`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `UserAccountInvited`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `TaskAssigned`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `AuthTokens`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

ALTER TABLE
    `TaskComments`
    MODIFY
    COLUMN `id` BIGINT AUTO_INCREMENT;

-- FOREIGN KEY
ALTER TABLE
    `Users`
    ADD
        FOREIGN KEY (`role_code`) REFERENCES `Roles` (`code`);

ALTER TABLE
    `Projects`
    ADD
        FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

ALTER TABLE
    `Tasks`
    ADD
        FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE
    `Tasks`
    ADD
        FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

ALTER TABLE
    `TaskAttachments`
    ADD
        FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE
    `TaskAttachments`
    ADD
        FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

ALTER TABLE
    `ProjectTags`
    ADD
        FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE
    `ProjectTags`
    ADD
        FOREIGN KEY (`tag_id`) REFERENCES `Tags` (`id`);

ALTER TABLE
    `TaskTags`
    ADD
        FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE
    `TaskTags`
    ADD
        FOREIGN KEY (`tag_id`) REFERENCES `Tags` (`id`);

ALTER TABLE
    `Notifications`
    ADD
        FOREIGN KEY (`caused_by`) REFERENCES `Users` (`id`);

ALTER TABLE
    `NotificationNotified`
    ADD
        FOREIGN KEY (`notification_id`) REFERENCES `Notifications` (`id`);

ALTER TABLE
    `NotificationNotified`
    ADD
        FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE
    `ProjectMembers`
    ADD
        FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE
    `ProjectMembers`
    ADD
        FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE
    `ProjectMembers`
    ADD
        FOREIGN KEY (`role_code`) REFERENCES `Roles` (`code`);

ALTER TABLE
    `ProjectMemberInvited`
    ADD
        FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE
    `ProjectMemberInvited`
    ADD
        FOREIGN KEY (`user_account_invited_id`) REFERENCES `UserAccountInvited` (`id`);

ALTER TABLE
    `TaskAssigned`
    ADD
        FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE
    `TaskAssigned`
    ADD
        FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE
    `AuthTokens`
    ADD
        FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE
    `TaskComments`
    ADD
        FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE
    `TaskComments`
    ADD
        FOREIGN KEY (`parent_comment_id`) REFERENCES `TaskComments` (`id`);

ALTER TABLE
    `TaskComments`
    ADD
        FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

-- CREATE INDEX
-- MYSQL
-- INSERT SAMPLE DATA FOR ALL TABLES 
-- 1. Roles
INSERT INTO `Roles` (`code`, `name`)
VALUES (1, 'Admin');

INSERT INTO `Roles` (`code`, `name`)
VALUES (2, 'Member');

INSERT INTO `Roles` (`code`, `name`)
VALUES (3, 'Owner');

INSERT INTO `Roles` (`code`, `name`)
VALUES (4, 'Observer');

-- 2. Users
INSERT INTO `Users` (`id`,
                     `username`,
                     `email`,
                     `full_name`,
                     `role_code`,
                     `status`,
                     `password_hash`,
                     `salt`)
VALUES (1,
        'peter01',
        'peter01@yopmail.com',
        'Peter',
        1,
        1,
        '20194891  02qwoioqw',
        '15151512');

INSERT INTO `Users` (`id`,
                     `username`,
                     `email`,
                     `full_name`,
                     `role_code`,
                     `status`,
                     `password_hash`,
                     `salt`)
VALUES (2,
        'john01',
        'john01@yopmail.com`,`John',
        'Johnathan',
        2,
        1,
        '02914qwjiooq',
        '2314u1241');

INSERT INTO `Users` (`id`,
                     `username`,
                     `email`,
                     `full_name`,
                     `role_code`,
                     `status`,
                     `password_hash`,
                     `salt`)
VALUES (3,
        'jane01',
        'janefoster09@yopmail.com',
        'Jane',
        2,
        1,
        '10941902qwiour',
        '8720984eu');

-- 3. Projects
INSERT INTO `Projects` (`name`,
                        `description`,
                        `status`,
                        `thumbnail_url`,
                        `priority`,
                        `privacy`,
                        `created_by`,
                        `deadline`,
                        `started_at`,
                        `thumbnail`)
VALUES ('Project 1',
        'This is project 1',
        1,
        'img1.url',
        '1',
        '1',
        '01',
        '12/05/2024',
        '05/02/2024',
        'img1.jpg');

-- 4. ProjectMembers
INSERT INTO `ProjectMembers` (`project_id`, `user_id`, `role_code`)
VALUES (1, 1, 3);

INSERT INTO `ProjectMembers` (`project_id`, `user_id`, `role_code`)
VALUES (1, 2, 2);

-- 5. Tasks
INSERT INTO `Tasks` (`parent_task_id`,
                     `project_id`,
                     `status`,
                     `created_by`,
                     `name`,
                     `description`,
                     `due_date`,
                     `started_at`)
VALUES (NULL,
        1,
        1,
        '01',
        'Task 1',
        'This is task 1',
        '2024-02-12',
        '2024-02-05')
    INSERT
INTO
    `Tasks` (`parent_task_id`,
             `project_id`,
             `status`,
             `created_by`,
             `name`,
             `description`,
             `due_date`,
             `started_at`)
VALUES
    (
    1, 2, 1, '02', 'Task 2', 'This is task 2', '2024-02-20', '2024-02-10'
    )