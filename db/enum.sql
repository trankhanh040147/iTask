
enum ActionType {
  TEAM_ADDED
  TEAM_REMOVED
  PROJECT_ADDED
  PROJECT_REMOVED
  PROJECT_ARCHIVED
  DUE_DATE_ADDED
  DUE_DATE_REMOVED
  DUE_DATE_CHANGED
  DUE_DATE_REMINDER
  TASK_ASSIGNED
  TASK_MOVED
  TASK_ARCHIVED
  TASK_ATTACHMENT_UPLOADED
  COMMENT_MENTIONED
  COMMENT_OTHER
}

enum NotificationFilter {
  ALL
  UNREAD
  ASSIGNED
  MENTIONED
}

enum RoleCode { 
  owner
  admin
  member
  observer
}

enum ShareStatus {
  INVITED
  JOINED
}

enum RoleLevel {
  ADMIN
  MEMBER
}

enum ActionLevel {
  ORG
  TEAM
  PROJECT
}

enum ObjectType {
  ORG
  TEAM
  PROJECT
  TASK
  TASK_GROUP
  TASK_CHECKLIST
  TASK_CHECKLIST_ITEM
}


enum MyTasksStatus {
  ALL
  INCOMPLETE
  COMPLETE_ALL
  COMPLETE_TODAY
  COMPLETE_YESTERDAY
  COMPLETE_ONE_WEEK
  COMPLETE_TWO_WEEK
  COMPLETE_THREE_WEEK
}

enum MyTasksSort {
  NONE
  PROJECT
  DUE_DATE
}

enum ActivityType {
  TASK_ADDED
  TASK_MOVED
  TASK_MARKED_COMPLETE
  TASK_MARKED_INCOMPLETE
  TASK_DUE_DATE_CHANGED
  TASK_DUE_DATE_ADDED
  TASK_DUE_DATE_REMOVED
  TASK_CHECKLIST_CHANGED
  TASK_CHECKLIST_ADDED
  TASK_CHECKLIST_REMOVED
}