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

# Standup Configuration

┌─ General Configuration ────────────────────────────────────┐
| ┌─ Name ───┐ ┌─ Default Section ─┐ ┌─ Sections Per Row ─┐  |
| | John Doe | | Worked On         | | 2                  |  |
| └──────────┘ └───────────────────┘ └────────────────────┘  |
| ┌─ Start of Week Day ─┐ ┌─ Holidays ──────┐                |
| | Monday              | | Saturday,Sunday |                |
| └─────────────────────┘ └─────────────────┘                |
|                                                            |
└────────────────────────────────────────────────────────────┘
┌─ Standup Sections ──────────────────────────────────────────────────┐
| ┌─────────────────────────────────────────────────────────────────┐ |
| | Full Section Name | Short | Description                         | |
| | Worked On         | wo    | Tasks worked on for the day         | |
| | Blockers          | bl    | Blockers affect completion of tasks | |
| | At Risk           | ar    | May not complete due to some issue  | |
| | PR Reviews        | prr   | Reviews of pull requests            | |
| └─────────────────────────────────────────────────────────────────┘ |
└─────────────────────────────────────────────────────────────────────┘
```

### Adding a task 

#### to default section

```
❯ standup task add -d "This is a sample task that says i did something"
   SUCCESS  Added task with id c4dca346-2c24-4532-a5a2-ad5e14b7f2d4 to standup
```

#### to specific section

```
❯ standup task add -s "PR Reviews" -d "This is a another sample task that says i did something"
 SUCCESS  Added task with id bc341e01-b22e-482f-ba19-2ace41e7d4ab to standup
```

### to specific dates

**NOTE**: date --day, --month (eg January) and --year options are individual and optional and take whatever is today by default

```
❯ standup task add -s "PR Reviews" -d "This is a sample task that says i did something another day" --day 26
 SUCCESS  Added task with id 8d740441-03d0-4466-97b2-2c0c81335263 to standup
```

### get standup report

```
❯ standup report

# Standup information

 Name  John Doe
 Type  Specific Day

# Report

┌─ Standup Date 27 April 2023 Thursday IST ────────────────────────────────────────────────────────────────────────────┐
| ┌─ PR Reviews ──────────────────────────────────────────────┐ ┌─ Worked On ───────────────────────────────────────┐  |
| |  Desc  Reviews of pull requests                           | |  Desc  Tasks worked on for the day                |  |
| |                                                           | |                                                   |  |
| | • This is a another sample task that says i did something | | • This is a sample task that says i did something |  |
| └───────────────────────────────────────────────────────────┘ └───────────────────────────────────────────────────┘  |
|                                                                                                                      |
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```

#### For the entire week

```
❯ standup report -w

# Standup information

 Name  John Doe
 Type  Weekly

# Reports

Name: John Doe
┌─ Standup Date 24 April 2023 Monday IST ─┐
|  INFO  No Standup recorded, skipping    |
|                                         |
└─────────────────────────────────────────┘

┌─ Standup Date 25 April 2023 Tuesday IST ─┐
|  INFO  No Standup recorded, skipping     |
|                                          |
└──────────────────────────────────────────┘

┌─ Standup Date 26 April 2023 Wednesday IST ─────────────────────────┐
| ┌─ PR Reviews ──────────────────────────────────────────────────┐  |
| |  Desc  Reviews of pull requests                               |  |
| |                                                               |  |
| | • This is a sample task that says i did something another day |  |
| └───────────────────────────────────────────────────────────────┘  |
|                                                                    |
└────────────────────────────────────────────────────────────────────┘

┌─ Standup Date 27 April 2023 Thursday IST ────────────────────────────────────────────────────────────────────────────┐
| ┌─ PR Reviews ──────────────────────────────────────────────┐ ┌─ Worked On ───────────────────────────────────────┐  |
| |  Desc  Reviews of pull requests                           | |  Desc  Tasks worked on for the day                |  |
| |                                                           | |                                                   |  |
| | • This is a another sample task that says i did something | | • This is a sample task that says i did something |  |
| └───────────────────────────────────────────────────────────┘ └───────────────────────────────────────────────────┘  |
|                                                                                                                      |
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

┌─ Standup Date 28 April 2023 Friday IST ─┐
|  INFO  No Standup recorded, skipping    |
|                                         |
└─────────────────────────────────────────┘

 INFO  week still in progress/exceeded today


```

#### For specific date

```
❯ standup report --day 26

# Standup information

 Name  John Doe
 Type  Specific Day

# Report

┌─ Standup Date 26 April 2023 Wednesday IST ─────────────────────────┐
| ┌─ PR Reviews ──────────────────────────────────────────────────┐  |
| |  Desc  Reviews of pull requests                               |  |
| |                                                               |  |
| | • This is a sample task that says i did something another day |  |
| └───────────────────────────────────────────────────────────────┘  |
|                                                                    |
└────────────────────────────────────────────────────────────────────┘
```
