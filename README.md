# Standup

This is a poc app for recording and reporting standups

## Building
First have golang which supports go mod
```
❯ scripts/build.sh
❯ ./standup blah
```

## Simple commands

### View standup config

```
❯ standup config view
Configuration Parameters: 

Sections:
	At Risk: Possible non-completion due to various reasons
	Blockers: Blockers for completing tasks that are affecting completion
	PR Reviews: All pull request reviews
	Worked On: What tasks were worked on for the day
Default Section: Worked On
Start of weekday:  Monday
Holidays:
	-Saturday
	-Sunday
```

### Adding a task 

#### to default section

```
❯ standup task add -d "description"  
```

#### to specific section

```
❯ standup task add -s "section name" -d "description"
```

### to specific dates

**NOTE**: date options are individual and optional and take whatever is today by default

```
❯ standup task add -d "description" --day 10 --month January --year 2020
```

### get standup report

```
❯ standup report
```

#### For the entire week

```
❯ standup report -w
```

#### For specific date

```
❯ standup report --day 10 --month January --year 2020
```
