# Pipestack - SMART Task Manager

## Overview

Have you read "The One Minue Manager" ?

It's a fantastic book, short, succinct and packed full of excellent suggestions about how to manage workloads, get things done and work effectively with others.

The biggest problem that I have come across in my day to day working life, and this is without exception;

In effective, badly executed DELEGATION. It's as though people don't even know the meaning of the word! This simple application sets out to resolve that issue. It builds upon a wonderful cliche
"Doing Things Right, and Doing the Right Things"

I call this "The DTR Method", which I attribute to the late, great "Ann Rutherford". Here's how it works

Managers are called upon to ensure that their teams are "Doing the Right Things", to facilitate this effectively however, the manager has to be responsible for "Doing Things Right".
This is crucially important when it comes to the art of delegation. Far too often I have come across managers that delegate with simple verbal communication, or worse still by email!!

The first problem those managers run into, is Who is doing what ? and by When is it to be done ?

So Pipestack sets up a Task delegation and management mechanism that tracks who has been asked to do what, and by when. It provides a novel approach to prioritisation by using the, Computer Science, concepts of :- 
 
 * The Stack
 * The Pipe
 * The Heap
 
 *Ooo! Freaky Geeky*
 
 Everything is built upon the Task, which is the single atomic unit of the application. A Task is a model of a real world task, and it implements the following attributes 
 
 id, name, parent_id, worklog_id, team_id, description, assignee, creator, startdate, duedate, createddate, modifieddate, status(New, Open, Closed, Paused, Cancelled ), description, estimatedeffort, actualeffort, prioritychannel ( Stack, Pipe, Heap)
 
 Every task can have a relationship with a parent task, it's own worklog and a team
 
 ## Tasks
 
 Tasks can only be re-assigned whilst in a New, or Paused state.
 A User can only have ONE Task open at any moment, and this is displayed in the Currently Working On box of their workboard
 When ever a user either changes the Task Status, or if the try to Open another task (when one is already open) a modal popup appears to record details of their time.
 
 ## Users
 
 Their are two types of User
 
  * Managers
  * Workers
  
 Managers signup via the main application, and create the initial Team. 
 
 Users are invited to the Team via email. Where they're an existing user of Pipestack we automatically add them to the new team
 
 ## Teams
 
 Teams are created when a new user (Manager) signs up via the application front end. By default a Manager is associated with 1 Team. However, users ( both managers and workers )
 can create new teams. The Workboards utilise the Teams resource to build a layout.
 
 An admin panel is provided for each team where configurations for the behaviour of the team can be applied
 
  * Auto-assignment (Least busy, Round Robin, Specific User, Off)
  * Minimum days ahead of start date that a Task can assigned to The Heap. 
 
 ## Task prioritisation
 When a new Task is created the Manager must decide upon the priority, based on the paradigms set out below. This is presented in a step through wizard format. The Manager first chooses the priority. Then a resource infographic is displayed for each team member, showing the size of their Stack, Pipe Queue. The manager can then consider which resource to assign the task to.
 Notice that their is no ability for user assignment from the Heap ( see The Heap below).
 
   
 ## The Stack
 The Stack is the HIGH priority task queue. Each user has one Stack, and they can only take the Task that is currently on top of the Stack. Other Stack Tasks are immutable until the Top Stack task has been Closed, or assigned elsewhere. Only the Team Manager has the authority to re-order a users Stack.
 
  
 ## The Pipe
 The Pipe is a pipeline of Tasks, which a user can pick and choose how and when they work on them. Consider these the MEDIUM priority tasks
  
 ## The Heap
 
 The Heap is a LOW priority task pile, and it is ONLY visible to the Team Manager. Tasks added to The Heap must still have ALL their details fulfilled, and must be SMART.
 Tasks added to the Heap must have a future start date that is at least (x) number of days in the future. The value of X is configured in the Team Admin panel.
 
 The Heap has an auto-assign function, which captures Tasks by start date, where the start date is within (Z) days of today. It then applies the auto assign logic and assigns the task into The Pipe for a user.
 
 
 ## Workboard
 
 Workboards provide a Team selection function. This is to enable a user to hold a single account and yet be associated to more then one team. In such cases the user will be required to meet with the manage of each team for a 1-2-1 once a week.
 
 A default workboard is setup for Workers, which shows a workboard widget for each of the following :-
 
  * Open Tasks on the Stack
  * New Tasks on the Stack
  * Open Tasks in the Pipe
  * New Tasks in the Pipe
  * Overdue Open Tasks on the Stack
  * Overdue Open Tasks in the Pipe
  * Closed Tasks from the last 7 days (i.e the last week)
  
  In addition the workboard is user configurable. Users can add / remove widgets, configure the filters for each widget (see widgets). They can drag and drop the widgets into different positions on the workboard.
  
  Thus the Workboard needs to implement some form of dashboard, drag n drop configuration system, that should provide a grid layout. The number of columns available on the Grid should also be configurable, probably between 1 and 5
  makes sense.
  
  ## Widgets
  
  Widgets provide a basic Workboard component for displaying information about a Task. It provides a scrollable window populated with rows of columns.
  The columns visible can be configured to be shown or hidden, with the exception of the description field which often has too much detail in it.
  The description field is instead shown, as an icon that when clicked presents a popup modal with the description info in.
  Each row represents a single task, and each row is colour coded ( To be decided but maybe something like)
  
   * White - New (In Tolerance)
   * Yellow - New ( Past Start date)
   * Orange - New ( 24hours until Due date)
   * Red - !Closed ( Past due date)
   * Blue - Deferred
   * Gray - Cancelled
 
 ## One-to-One Board
 
 ## Scrum Board
 
 ## Projects
 
 ## Time Recording
 
 Each Task has a related Worklog, which records the user, start_datetime, and stop_datetime, activity, total_time. This is driven by a popup modal which is activated whenever a Task status changes from Open to some other value.
 
 
 ## Configuration
 
 ## Online Help System
 
 ## Sign up wizard
 
 SMART Task management application, that takes the sting out of organising and managing a team of people.
 This journey is how users first sign up, and is designed to on-board users as simply as possible, and each of the steps below is a page in it's own right
 
 ### Initial Sign up journey
 
   * (Username, Password) Social Login, Phone ( Mobile )
   * Team Name
   * How many people in the team
   * What does your team do ?
   * About you team - Name, Email, Role
   * 1-2-1 Setup - Choose times, and assign to 
   
   The about your team page, renders a form which which expects a Name, Email and Role for each member of the team, it renders form rows equal to the number of team members identified in the previous question.
   
   The 1-2-1 Setup feature, provides a time (24:00 UTC ** Fixme) and day ( Full Week Mon - Sun) drop down, along with a drop down chooser for each team member. This sets up a weekly reminder ( 30 minutes ahead)for each team member. This will be the day and time that they meet with their manager for a short 1-2-1 (circa 10 minutes)
 
 ## Social Login
 
 Social Login utilises a responsive grid panel which displays icons for ALL the various login authentication types. Social Login uses Buffalos 'Gothic' package to take care the OAuth stuff.
 and retrieves the users, Firstname, Lastname, NickName, Name, Email, and Profile picture and creates a Pipestack users 
 
 ## Reports
 
 ## Zapier integration
 
 ## Telegram integration
 
 ## Pipestack API
 
 # Use Cases
 
 ## 
 
 