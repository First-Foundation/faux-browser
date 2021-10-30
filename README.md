# faux-browser

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/First-Foundation/faux-browser/agent-terminal-go-build?style=plastic)
![GitHub last commit](https://img.shields.io/github/last-commit/First-Foundation/faux-browser?style=plastic)
![GitHub](https://img.shields.io/github/license/First-Foundation/faux-browser?style=plastic)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/First-Foundation/faux-browser?style=plastic)

**Fake Automated User eXperience for generating user browser activity.**

The purpose of faux-browser is to generate more realistic traffic generation for blue/purple-team exercise ranges or labs.

## Quick Start
- Run `go get golang.org/x/net/html`
- Make browser changes as desired (see below).
- Make profile changes as desired (see below).
- Search for `USER TODO` and make changes as you desire.
- Run `./build.sh`
- Copy the appropriate binary to your machine(s), renaming it to the browser you want to emulate.
  - For Windows, try `chrome.exe`, `edge.exe`, `firefox.exe`, or `iexplore.exe`
  - For Linux, try `chromium` or `firefox`
  - For MacOS ("darwin"), try `Google Chrome.app` or `Safari.app`

## Profiles

Most profiles mimic a certain type of user, with the exceptions being `profile_common` and `profile_testing`.
  - `profile_common` is a list of common sites/searches that all profiles will conduct.
  - `profile_testing` should only be used in a testing scenario, as it ignores schedule hours.

The purpose of a profile is to have an additional subset of searches/sites that a certain type of user would perform (for example, a member of the sales team will not be going to all of the same sites that a member of the cybersecurity team would be).

Further, each profile has different schedules. Schedules have different start/stop times for the business day, take lunch breaks at different times; currently, all schedules (except for `profile_testing`) do not browse on the weekends.

A short breakdown of the different type of employee profiles are below:
  - `profile_cybersecurity` - This profile mimics an employee either a part of the Cybersecurity team, or an aspiring ethical hacker. Sites visited may be suspicious (such as exploit sites, and searching for ransomware strains), but the "intent" of the searches are not malicious. This profile in particular may serve as a red-herring for exercises.
  - `profile_developer` - This profile mimics programmers and scripters, members of the development team, etc. Likely not accurately mimicking since it won't visit StackOverflow near as often as a real developer would.
  - `profile_employee_average` - This profile is intended to be the average, non-driven employee. They're in the rat race, showing up to work, doing what they need to do but maybe not much more.
  - `profile_employee_driven` - This profile is for the driven employee, who is driven to do their job and has a "stronger" work ethic. This profile includes searches for how to perform better at the current job *and* looking for other jobs if they aren't happy with their current workplace.
  - `profile_gamer` - This profile is for the employee who puts games over work. Even during work hours, they want to keep up with gaming news and communities.
  - `profile_it` - Mimicking a member of I.T., this profile mimics system and network administrators. Not quite as much of a red-herring as the cybersecurity profile, they may still conduct some searches that are close.
  - `profile_management` - CEOs, middle management, and brown-nosers are those mimicked by this profile. They'll show up late and leave early, take longer lunches. Note: If you need a profile to visit a malicious site internal to your exercise range, it wouldn't be a bad idea to modify this profile to include a visit.
  - `profile_newsaddict` - The employees mimicked by this profile need to keep up-to-date on the state of things going on while they work.
  - `profile_redditor` - They visit reddit at work, 'nough said.
  - `profile_researcher` - This profile is for scientists and researchers, interested in mathematics, data science, and such.
  - `profile_sales` - Sales team members are mimicked by this profile.
  - `profile_threat` - This profile mimics the employee that may (or may not) be a threat. Some of the sites they visit may be pretty suspicious...

## Selecting specific profiles

Profiles are selected at random, with some (fake) weighted math thrown in. This allows you to add the executable to an image and spin up 10s of 100s of instances and it will simulate multiple profile types naturally.

If you want a specific machine to use a specific profile, you can either modify the source code in the main function of `faux-browser.go` and compile an executable for that particular machine, *or* you can set the `FAUX_BROWSER_ROLE` environment variable to one of the following values:
| **Value** | **Profile** |
| --- | --- |
| `cybersec` | The `profile_cybersecurity` profile. |
| `dev` | The `profile_developer` profile. |
| `driven` | The `profile_employee_driven` profile. |
| `employee` | The `profile_employee_average` profile. |
| `gamer` | The `profile_gamer` profile. |
| `manage` | The `profile_management` profile. |
| `news` | The `profile_news` profile. |
| `reddit` | The `profile_redditor` profile. |
| `research` | The `profile_researcher` profile. |
| `sales` | The `profile_sales` profile. |
| `tech` | The `profile_it` profile. |
| `testing` | The testing profile, not recommended for a live environment. |
| `threat` | The `profile_threat` profile. |

*PLEASE NOTE: If both the `FAUX_BROWSER_ROLE` is set and the main function is modified to call a specific profile, and the values are different, only one of these profiles will be used!*

## Adding Sites and Searches

You may need to add Sites and Searches to suit your needs (especially internal sites for the exercise range/lab). Adding sites/searches comes in a few different ways:
- If you are introducing domains that you want visited, you may want to add them to the allow list in `browser.go`. There is a random chance that any profile will visit a domain from this overarching list.
- If you want *any* profile to visit a URL or conduct a search, drop it in `profile_common.go`.
- If you want only a particular profile to visit the site or conduct a search, drop it in the appropriate profile.
