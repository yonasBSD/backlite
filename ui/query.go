package ui

const itemLimit = 25

// TODO no need to select task field for lists, need claimed_at for running tasks

const selectRunningTasks = `
	SELECT 
	    id,
	    queue,
	    null,
	    attempts,
	    wait_until,
	    created_at
	FROM 
	    backlite_tasks
	WHERE
	    claimed_at IS NOT NULL
	LIMIT ?
`
const selectCompletedTasks = `
	SELECT
	    id,
		created_at,
		queue text,
		last_executed_at,
		attempts,
		last_duration_micro,
		succeeded,
		null,
		expires_at,
		error
	FROM
	    backlite_tasks_completed 
	WHERE
	    succeeded = ?
	ORDER BY
	    last_executed_at DESC
	LIMIT ?
`
