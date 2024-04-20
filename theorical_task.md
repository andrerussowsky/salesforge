To design a system that can handle sending emails in equal intervals between sending hours while considering step wait days and daily mailbox capacities, we need to consider a few key components and their interactions. Here's an overview of how such a system could work:

# Components:
- API Server: Handles incoming requests to create sequences, update sequence steps, and manage mailbox and contact information.
- Scheduler: Responsible for scheduling and triggering email sending tasks at specified intervals.
- Email Sender: Component responsible for sending emails based on the scheduled tasks and the availability of mailboxes' daily capacities.
- Database: Stores information about sequences, steps, mailboxes, contacts, and the status of email sending tasks.

# Workflow:
1. Sequence Creation and Configuration:
- Users create sequences and configure the sending hours, steps, associated mailboxes, and contact lists through the API server.
- Each sequence can have multiple steps, each with its own email content and wait days between sends.
- Sequences are associated with one or more mailboxes, each with a maximum daily sending capacity.

2. Scheduling Email Sending Tasks:
- The scheduler component periodically checks the database for scheduled email sending tasks.
- For each sequence, the scheduler calculates the next email sending time based on the sending hours and step wait days.
- It schedules email sending tasks at equal intervals between the specified sending hours, taking into account the capacities of the associated mailboxes.

3. Email Sending Process:
- When an email sending task is triggered, the email sender component retrieves the sequence details from the database.
- It checks the daily sending capacity of each associated mailbox and distributes the email sending workload evenly among them.
- The email sender processes the contacts associated with the sequence, sends emails to them in batches, and updates the database with the sent status.

4. Retries and Error Handling:
- If an email sending task fails due to network issues or other errors, the system retries the task based on a predefined retry policy.
- Failed email sending tasks are logged and monitored for analysis and troubleshooting.

# Horizontal Scaling:
- To handle increased load and ensure high availability, the system can be horizontally scaled by deploying multiple instances of the API server, scheduler, and email sender components.
- Load balancers can distribute incoming requests across the instances of the API server.
- Each instance of the scheduler can operate independently, distributing the workload evenly across the available resources.

By following this approach, the system can efficiently send emails in equal intervals between sending hours while considering step wait days and daily mailbox capacities. Horizontal scaling ensures that the system can handle increased load and retries provide resilience in case of failures. Additionally, proper monitoring and logging help in identifying and resolving issues quickly.