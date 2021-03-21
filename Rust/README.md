# Getting Started with Rust on Windows and VSCode

## install rust

https://www.rust-lang.org/tools/install click the gargantuan button labeled “Rustup-Init.exe.”

## verify rust is installed

Open a new PowerShell window or command prompt and execute the following:

```
rustup --version
```

## configure visual studio code

There are a variety of VS Code extensions that support Rust. **Rust Extension Pack** Not only does it bundle the Rust RLS official extension, but it includes the most popular cargo and TOML plugins available. To install this extension pack:

- Navigate to the Extensions panel (or type Ctrl-Shift-X).
- Enter “Rust Extension Pack” in the search panel.
- Click the Install button.

## create a simple HELLO WORLD project

```
> cargo new hellorust
```

https://doc.rust-lang.org/nightly/edition-guide/editions/creating-a-new-project.html

## create a build task

Open the command palette (Ctrl-Shift-P) and type in “build” and select “Tasks: Configure Default Build Task”

Select “Rust: cargo build”

This should create a tasks.json file that has a single, default build task.

```
{
	"version": "2.0.0",
	"tasks": [
		{
			"type": "cargo",
			"subcommand": "build",
			"problemMatcher": [
				"$rustc"
			],
			"group": {
				"kind": "build",
				"isDefault": true
			},
			"label": "Rust: cargo build - hellorust"
		},
		{
			"type": "cargo",
			"subcommand": "test",
			"problemMatcher": [
				"$rustc"
			],
			"group": "test",
			"label": "Rust: cargo test - hellorust"
		}
	]
}
```

## configure debugging configuration

1. Install the C/C++ extension.
2. Install the Native Debug extension.
3. In the Debug panel, add a new configuration.
4. Specify the “C++ (Windows)” environment. This will create a new launch.json file.
5. This will create a new launch.json file. Change the program value to be “\${workspaceFolder}/hellorust/target/debug/hellorust.exe” and save the file.

Your launch.json should look similar to mine:

```
{
    // IntelliSense を使用して利用可能な属性を学べます。
    // 既存の属性の説明をホバーして表示します。
    // 詳細情報は次を確認してください: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "(Windows) 起動",
            "type": "cppvsdbg",
            "request": "launch",
            "program": "${workspaceFolder}/hellorust/target/debug/hellorust.exe",
            "args": [],
            "stopAtEntry": false,
            "cwd": "${workspaceFolder}",
            "environment": [],
            "externalConsole": false
        }
    ]
}
```

Now, when you hit F5, VS Code will start debugging a new instance of your program. You can add breakpoints where needed.

## run rust online

https://play.rust-lang.org/
