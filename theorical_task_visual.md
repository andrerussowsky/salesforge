```
                             +-------------------+
                             |     API Server    |
                             +-------------------+
                                      |
                                      |
                                      v
                          +-----------------------+
                          |        Database       |
                          +-----------------------+
                                      |
                                      |
                                      v
                     +----------------------------+
                     |        Scheduler           |
                     +----------------------------+
                              |       |
                 +------------+       +------------+
                 v                                 v
      +----------------------+    +----------------------+
      |      Email Sender    |    |      Email Sender    |
      +----------------------+    +----------------------+
```

In this diagram:
- The API Server handles incoming requests to manage sequences, steps, mailboxes, and contacts.
- The Database stores information about sequences, steps, mailboxes, contacts, and scheduled email sending tasks.
- The Scheduler component periodically schedules email sending tasks based on sequence configurations and mailbox capacities.
- Email Sender components retrieve scheduled tasks from the database, distribute email sending workload among available mailboxes, and send emails to contacts.
- Horizontal scaling is achieved by deploying multiple instances of the API Server, Scheduler, and Email Sender components to handle increased load and ensure high availability.
