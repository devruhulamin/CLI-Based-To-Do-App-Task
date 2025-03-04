# Todo CLI App

The app uses an environment variable `TODO_FILENAME` to specify the file where tasks will be saved. You can set it like this:

```bash
export TODO_FILENAME=tasks.json
```

### Commands and Flags

- **Add a new task:**

```bash
./todo -add "Your task here"
```

- **Delete a task (by index):**

```bash
./todo -del 1
```

- **List all tasks:**

```bash
./todo -list
```

- **Mark a task as complete (by index):**

```bash
./todo -complete 1
```

- **Show details of all tasks:**

```bash
./todo -det
```

- **Show only pending (uncompleted) tasks:**

```bash
./todo -pending
```

- **Get a specific task (by index):**

```bash
./todo -i 1
```

