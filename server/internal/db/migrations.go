package db

import (
	"log"
)

// RunMigrations runs the database migrations
func RunMigrations() error {
	schema := `
	-- Create users table if not exists
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		rating INT DEFAULT 0,
		tasks_solved INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- Create index on email for faster lookups
	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

	-- Create index on rating for leaderboard queries
	CREATE INDEX IF NOT EXISTS idx_users_rating ON users(rating DESC);

	-- Create tasks table if not exists
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		difficulty VARCHAR(20) NOT NULL CHECK (difficulty IN ('easy', 'medium', 'hard')),
		constraints TEXT,
		starter_code_go TEXT,
		starter_code_js TEXT,
		starter_code_python TEXT,
		starter_code_cpp TEXT,
		starter_code_java TEXT,
		time_limit_ms INT DEFAULT 1000,
		memory_limit_kb INT DEFAULT 131072,
		is_community BOOLEAN DEFAULT FALSE,
		solved_percent DECIMAL(5,2) DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- Create index on difficulty for filtering
	CREATE INDEX IF NOT EXISTS idx_tasks_difficulty ON tasks(difficulty);

	-- Create index on is_community for filtering community tasks
	CREATE INDEX IF NOT EXISTS idx_tasks_is_community ON tasks(is_community);

	-- Create test_cases table if not exists
	CREATE TABLE IF NOT EXISTS test_cases (
		id SERIAL PRIMARY KEY,
		task_id INT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
		input TEXT NOT NULL,
		expected TEXT NOT NULL,
		is_hidden BOOLEAN DEFAULT FALSE
	);

	-- Create index on task_id for faster test case retrieval
	CREATE INDEX IF NOT EXISTS idx_test_cases_task_id ON test_cases(task_id);

	-- Create submissions table if not exists
	CREATE TABLE IF NOT EXISTS submissions (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		task_id INT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
		language VARCHAR(20) NOT NULL CHECK (language IN ('go', 'javascript', 'python', 'cpp', 'java')),
		code TEXT NOT NULL,
		status VARCHAR(50) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'running', 'accepted', 'wrong_answer', 'time_limit_exceeded', 'memory_limit_exceeded', 'runtime_error', 'compilation_error')),
		runtime_ms INT,
		memory_kb INT,
		passed_tests INT DEFAULT 0,
		total_tests INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	);

	-- Create index on user_id for user submission history
	CREATE INDEX IF NOT EXISTS idx_submissions_user_id ON submissions(user_id);

	-- Create index on task_id for task submission statistics
	CREATE INDEX IF NOT EXISTS idx_submissions_task_id ON submissions(task_id);

	-- Create composite index for user's submissions on a specific task
	CREATE INDEX IF NOT EXISTS idx_submissions_user_task ON submissions(user_id, task_id);

	-- Create index on status for filtering submissions
	CREATE INDEX IF NOT EXISTS idx_submissions_status ON submissions(status);

	-- Create index on created_at for recent submissions queries
	CREATE INDEX IF NOT EXISTS idx_submissions_created_at ON submissions(created_at DESC);
	`

	_, err := DB.Exec(schema)
	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}
