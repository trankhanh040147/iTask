CREATE TABLE `Users` (
  `id` BIGINT PRIMARY KEY,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `role_code` varchar(255) NOT NULL,
  `title` varchar(255),
  `status` int NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `salt` varchar(255) NOT NULL,
  `address` varchar(255),
  `phone` varchar(255),
  `dob` varchar(255),
  `profile_ava_url` varchar(255),
  `profile_cover_url` varchar(255),
  `is_email_verified` int,
  `bio` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `Projects` (
  `id` BIGINT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `description` varchar(255),
  `status` int NOT NULL,
  `thumbnail_url` varchar(255),
  `priority` int,
  `privacy` int,
  `created_by` bigint NOT NULL,
  `deadline` timestamp,
  `started_at` timestamp,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now()),
  `thumbnail` varchar(255)
);

CREATE TABLE `Tasks` (
  `id` BIGINT PRIMARY KEY,
  `parent_task_id` int,
  `project_id` bigint NOT NULL,
  `status` int NOT NULL,
  `created_by` bigint NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255),
  `position` float,
  `priority` int,
  `completed` bool NOT NULL,
  `due_date` timestamp,
  `started_at` timestamp,
  `completed_at` timestamp,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `TaskAttachments` (
  `id` BIGINT PRIMARY KEY,
  `task_id` int,
  `created_by` bigint,
  `downloads` int,
  `file_name` varchar(255),
  `file_path` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `Tags` (
  `id` BIGINT PRIMARY KEY,
  `tag_type` int,
  `name` varchar(255),
  `description` varchar(255),
  `position` float,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `ProjectTags` (
  `id` BIGINT PRIMARY KEY,
  `project_id` bigint,
  `tag_id` bigint,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `TaskTags` (
  `id` BIGINT PRIMARY KEY,
  `task_id` bigint,
  `tag_id` bigint,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `Notifications` (
  `id` BIGINT PRIMARY KEY,
  `caused_by` bigint,
  `action_type` varchar(255),
  `data` json,
  `created_on` timestamp DEFAULT (now())
);

CREATE TABLE `NotificationNotified` (
  `id` BIGINT PRIMARY KEY,
  `notification_id` bigint,
  `user_id` bigint,
  `read` bool,
  `read_at` timestamp DEFAULT (now())
);

CREATE TABLE `ProjectMembers` (
  `id` BIGINT PRIMARY KEY,
  `project_id` bigint,
  `user_id` bigint,
  `added_at` timestamp DEFAULT (now()),
  `role_code` varchar(255)
);

CREATE TABLE `ProjectMemberInvited` (
  `id` BIGINT PRIMARY KEY,
  `project_id` bigint,
  `user_account_invited_id` bigint
);

CREATE TABLE `UserAccountInvited` (
  `id` BIGINT PRIMARY KEY,
  `email` text NOT NULL,
  `invited_on` timestamp NOT NULL DEFAULT (now()),
  `has_joined` boolean NOT NULL DEFAULT false
);

CREATE TABLE `Roles` (
  `code` varchar(255) PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE `TaskAssigned` (
  `id` BIGINT PRIMARY KEY,
  `task_id` bigint,
  `user_id` bigint,
  `assigned_date` timestamp DEFAULT (now())
);

CREATE TABLE `AuthTokens` (
  `id` BIGINT PRIMARY KEY,
  `token_id` varchar(255) NOT NULL,
  `user_id` bigint NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `expires_at` timestamp NOT NULL
);

CREATE TABLE `TaskComments` (
  `id` BIGINT PRIMARY KEY,
  `task_id` bigint,
  `parent_comment_id` bigint,
  `created_by` bigint,
  `message` varchar(255),
  `pinned` bool,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

ALTER TABLE `Users` ADD FOREIGN KEY (`role_code`) REFERENCES `Roles` (`code`);

ALTER TABLE `Projects` ADD FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

ALTER TABLE `Tasks` ADD FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE `Tasks` ADD FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

ALTER TABLE `TaskAttachments` ADD FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE `TaskAttachments` ADD FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);

ALTER TABLE `ProjectTags` ADD FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE `ProjectTags` ADD FOREIGN KEY (`tag_id`) REFERENCES `Tags` (`id`);

ALTER TABLE `TaskTags` ADD FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE `TaskTags` ADD FOREIGN KEY (`tag_id`) REFERENCES `Tags` (`id`);

ALTER TABLE `Notifications` ADD FOREIGN KEY (`caused_by`) REFERENCES `Users` (`id`);

ALTER TABLE `NotificationNotified` ADD FOREIGN KEY (`notification_id`) REFERENCES `Notifications` (`id`);

ALTER TABLE `NotificationNotified` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `ProjectMembers` ADD FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE `ProjectMembers` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `ProjectMembers` ADD FOREIGN KEY (`role_code`) REFERENCES `Roles` (`code`);

ALTER TABLE `ProjectMemberInvited` ADD FOREIGN KEY (`project_id`) REFERENCES `Projects` (`id`);

ALTER TABLE `ProjectMemberInvited` ADD FOREIGN KEY (`user_account_invited_id`) REFERENCES `UserAccountInvited` (`id`);

ALTER TABLE `TaskAssigned` ADD FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE `TaskAssigned` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `AuthTokens` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `TaskComments` ADD FOREIGN KEY (`task_id`) REFERENCES `Tasks` (`id`);

ALTER TABLE `TaskComments` ADD FOREIGN KEY (`parent_comment_id`) REFERENCES `TaskComments` (`id`);

ALTER TABLE `TaskComments` ADD FOREIGN KEY (`created_by`) REFERENCES `Users` (`id`);


-- CREATE INDEX

-- INSERT SAMPLE DATA FOR ALL TABLES
-- 1. Roles
INSERT INTO `Roles` (`code`, `name`) VALUES ('01', 'Admin');
INSERT INTO `Roles` (`code`, `name`) VALUES ('02', 'Member');
INSERT INTO `Roles` (`code`, `name`) VALUES ('03', 'Owner');
INSERT INTO `Roles` (`code`, `name`) VALUES ('04', 'Observer');

-- 2. Users
INSERT INTO `Users` (`username`, `email`, `full_name`, `role_code`, `status`, `password_hash`, `salt`) 
VALUES ('peter01','peter01@yopmail.com','Peter','01',1,'$2a$10$','15151512');