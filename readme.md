# AES and Blowfish Encryption/Decryption Web Interface

## Overview
This project is a simple web application that provides an interface for encrypting and decrypting text using AES and Blowfish encryption algorithms. It has two main sections: encryption and decryption, allowing users to enter text, generate keys, and view encrypted or decrypted results. The backend server, written in Go, handles the encryption and decryption logic.

## Features
- Encrypt text using AES and Blowfish algorithms.
- Decrypt encrypted data back to original text.
- Generate AES and Blowfish keys of specified lengths.
- Copy encryption details to the decryption section, and vice versa.

## Technologies Used
- **Frontend:** HTML, CSS, JavaScript
- **Backend:** Go (Golang)
- **Libraries:**
  - Font Awesome for icons
  - Fetch API for communication between the frontend and backend

## Getting Started
### Prerequisites
- Go (Golang) installed on your local machine.
- A web browser to access the frontend.

### Running the Backend Server
1. Clone the repository.
2. Navigate to the project folder.
3. Run the Go server:
   ```sh
   go run main.go
   ```
   The server will start at `http://localhost:8080`.

### Running the Frontend
1. Open the `index.html` file in your browser.
2. The interface should be visible, allowing you to interact with the encryption and decryption features.

## Usage
### Encryption
1. Enter the text you want to encrypt in the "Text to Encrypt" field.
2. Select the desired AES key length (16, 24, or 32 bytes).
3. Generate an AES key or enter your own.
4. Enter the desired Blowfish key length (1-56 bytes) and generate or enter a Blowfish key.
5. Click "Encrypt" to get the encrypted result.

### Decryption
1. Enter the encrypted text in the "Text to Decrypt" field.
2. Provide the same AES and Blowfish keys used for encryption.
3. Click "Decrypt" to retrieve the original text.

### Copying Data
- Use the "Copy Encrypt Info to Decrypt" button to copy the encrypted text and keys from the encryption section to the decryption section.
- Use the "Copy Decrypt Info to Encrypt" button to do the opposite.

## Customization
### Styling
- The styling can be modified using the CSS styles provided in the `<style>` tag within the HTML file. Feel free to tweak the colors, font, and layout to suit your preferences.

### Key Length
- The key length for AES can be adjusted using the dropdown menu, while the Blowfish key length can be specified in the input field.

## Future Enhancements
- Add more encryption algorithms for flexibility.
- Implement user authentication for secure access.
- Enhance error handling and validation.

## License
This project is open-source and available under the [MIT License](LICENSE).

