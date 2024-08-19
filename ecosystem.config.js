module.exports = {
    apps: [
        {
            name: "app_name",
            script: "main.go",
            exec_interpreter: "bash",
            exec_mode: "fork",
            instances: 2,
            autorestart: true,
            watch: true,
            max_memory_restart: "1G",
            interpreter_args: "-c 'go run main.go agentd silent'"
        }
    ]
};
