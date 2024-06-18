# FXC Go Take-Home Technical Task

Welcome to the FXC Go Take-Home Technical Task! This repository contains a simple basis for a technical task that is loosely based on one of the kinds of technical solutions we must build at FXC.

If you've received this repo, then we'd like you to complete the task detailed below and submit it back to us. Please use the Go programming language to complete the task.

This is an opportunity for you to show your attention to detail and you knowledge of best practice, so please keep that in mind as you develop your solution.

### Brief

Frogo Baddins, a hobbit from the Shire, has a $10,000 and wants to send money to his friend in Philippines.
He has been gathering a database of the different FX rates offered by various FX payment companies, on different days, 
for different transfer amounts

His friend, Samrise, built for him a simple React UI to display a table of the following shape:

#### FX Rate

Date: 2024/01/10

| Transfer Amount   | Company 1 | Company 2 |
| ----------------- | --------  | --------- |
| 100               | 5.1       | 4.9       |
| 300               | 5.0       | 4.8       |
| 500               | 5.2       | 4.8       |


This table displays the FX rate for each Company, for each transfer amount, for a given day.

Frogo has hired you to build the backend microservice that will supply this frontend with data to populate the table. 

### Task

You are required to build a microservice that fulfil's Frogo's requirements, to supply Samrise's application with data.

The microservice should:
- Connect to the MariaDB instance
- Have an HTTP endpoint at an appropriately-named path, that takes the chosen date as a parameter
- Upon the endpoint being called, retrieve the relevant data from a table in MariaDB called `pricing`
- Transform the retrieved data into a form suitable for displaying in a table in Samrise's style (no transformation of the response should be required on the frontend)
- Return the transformed data in the response to the HTTP endpoint call

See [./sql/initialise.sql](./sql/initialise_data.sql) for more details about the database schema you will be working with.

### Evaluation Criteria

- The requirements given to you by Frogo aren't necessarily comprehensive. Where there is room for interpretation, use your expertise as a developer to choose a sensible solution
- Do your best to keep you code clean, simple and to make it extensible for future improvements
- Please include unit tests with your code. Unit test coverage does not have to be comprehensive - this code isn't really going to production - but please include enough tests that we can get a sense of your approach to testing
- Please use Git to version control your submission. Use any Git workflow you think is sensible and suits you

### Testing your solution

A Docker Compose file has been provided that should built and start your application in a docker container, as well as starting and 
populating a MariaDB instance. Please use this setup to ensure your application functions correctly.

### Confidentiality

This repository is confidential, please do not share it with anyone else, and delete it when you have finished with your application.
However, please be aware that you may be asked to extend this code if you get invited to a further interview, so don't delete it 
before then!

### Submitting your solution

Please return your solution, as a Git repository, in a compressed file, to your contact at FXC.

Best of luck, and happy coding!

The FXC Team
