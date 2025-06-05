# Take Home Challenge

## Introduction

This is a simulated legacy code base. Please pretend the portion you have been given is isolated from a larger code base
and was written by someone who:

* Is no longer at the company
* Did not have much experience when the code was written

Please spend less than 3 hours on this challenge. Anything you do not have time to do can be mentioned in the simulated
pull request.

**Use your best judgment so your changes are production-ready, reasonable for your ticket, and maintainable in the long
term.**

## Your ticket

Please update the `messages` controller to include an endpoint where a message is readable given its ID.

Additionally, please refactor the `messages` controller to be more production-ready and easier for future engineers to
maintain.

Scope examples:

* Adding new configuration for the `messages` controller does not affect other controllers. You can do this if necessary
* Changing how the service reads its existing configuration does affect other controllers and is out of scope for the
  ticket

Notes:

* You get to decide on the API design for the new read endpoint
* Your code changes should only affect the `messages` controller, any code paths taken by the `users` controller should
  be untouched
  * You can assume the given code (controllers, coordinators, etc.) is not reused elsewhere in the simulated legacy code
    base
* Take a holistic view of the `messages` controller for the scope of this ticket, configuration, controller,
  coordinator, tests, etc
* You can add, update, or remove 3rd party dependencies if it only affects the `messages` controller
* Pretend this is code currently running in production and your changes will go into production immediately after review
  * Do not make any breaking API changes, the frontend is currently using this code
* No metrics, alerting, or authentication required
* Please make `git` commits to the existing branch as you would normally at work
* Do not worry about adding too many commits, we often squash commits during merge
* No database migrations allowed
* You may not have enough time to do all the refactoring you would like
  * Do your best to demonstrate a well-rounded engineer given the time, leave the rest for the simulated PR suggestions

## Tests

To run the tests, ensure you have a recent version of Docker installed and run the below command in the root of the
project:

```
docker compose up
```

This will bring up a local database for testing. Run Go tests normally after containers are online.

To remove the Docker containers, run

```
docker compose down
```

## Pull request

Your changes will be reviewed by your peers. Please include a `pull-request.md` Markdown file to simulate a GitHub pull
request.

The file should have two headings.

The first heading should be suggestions to improve the service (all controllers). This is for suggestions that are
either outside the scope of the ticket or you did not have time to implement.

Please keep suggestions concise and include a short justification for each suggestion.

The second heading should be simulated pull request comments you are leaving for your peers before review starts.

Please only write comments for lines relevant to your ticket. Only comment on changes where you:

* Have added, updated, or removed a dependency (justify)
* Have a question you would like to answer before merging (assume the reviewers have never seen this section of the code before)
* Expect a change to merit a discussion during review

For comments, specify the code's location in the line above. The code's location is the path of the file relative to the
project root followed by ":L" and the line number.

Example review comment:

```
services/api/controllers/main.go:L1423
I replaced the existing dependency with [new name] for [reason].
```

## Submission

To submit, please:

* Create a `.zip` file containing the edited project
  * Make sure to include the `.git` directory but not `.idea` or `.vscode`
* Remember to include `pull-request.md`
* Send your `.zip` file to your recruiting point of contact
