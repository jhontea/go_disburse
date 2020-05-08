Suppose that you are an engineer at an online marketplace. In that marketplace, there are a lot of seller. Some of those seller, have sold their goods, and now they have some money in their balance. Of course they want to withdraw their money to their bank account. With a lot of seller who want to withdraw their money as quickly as possible, and they doing it on a sporadic time span, it will be problematic for the operational team on your company to disburse the money manually.

Fortunately, you come across a cool 3rd party service that can disburse your money to any recipient on any bank, without the pain of doing it manually. That 3rd party is called *slightly-big Flip*. *Slightly-big Flip* provide a simple API that you can use to disburse the money, so you're going to create a service for that.

Here's the requirement for your service:

- Your service will send the disbursement data to the 3rd party API
- Your service will then, save the detailed data about the disbursement from the 3rd party, in your local database
- Your service will also have the capability to check the disbursement status, and update the information on your database according to the response you get
    
You can learn about the API from the documentation [here](https://gist.github.com/luqmansungkar/9512940cac53f53bb4a024a1e5f70ef7)

## Information about the assignment

### Rules

- You should use PHP to work on this assignment
- You are not allowed to use any framework or any external libraries
- You must save the information you get from the API response, to your local database
- You must update the information you get from the disbursement status endpoint, to the related transaction in your local database
- The information that must be updated when you check the disbursement status are the following:
    - `status`
    - `receipt`
    - `time_served`
- You must use git when working on the project

### What will be scored

- Compliance to the rules, completeness, and correctness
- Use of OOP principle
- Project structure
- Code cleanliness
- Git commit

### Submission

- Provide a `README.md` that explain how to run the project. This should include something like how to run the disburse request, probably via command like `php disburse.php`, or any other way.
- Provide a config file as a place to provide the db connection and other credentials if any. This should be the only files that I need to edit to be able to run the project
- Provide a migration file that I can run to generate the database table. Just something simple, because the database should not be complex. Something that I can run with command like `php migration.php`

Send the repository public link by replying the previous email from me. The time to work on this assignment is 120 hours or 5 days since I send the instruction email to you. Remember to set the right visibility to the project repo so that I can also see the project.

Good luck!