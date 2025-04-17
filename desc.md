entities:
    Task:
        properties:
            ID          int       `json:"id"`
            Title       string    `json:"title"`
            Description string    `json:"description"`
            CreatedAt   time.Time `json:"created_at"`
            Done        bool      `json:"done"`
    Assignee:
        properties:
            ID    int
            Name  string
            Email string


use-case:
    Task:
        - user can create a task and assigne to assignee by id 
        - user can change the task status to done 
        - user can delete a task 
        - user can list all the task 
        - user can list all done tasks

    Assignee:
        - user can create an assignee
        - user can delete an assignee 
        - user can edit an assignee 
        - user can list assignee
 
