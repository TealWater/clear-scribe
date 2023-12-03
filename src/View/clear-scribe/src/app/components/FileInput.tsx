
"use client"

import { useState } from 'react';

export default function FileInput() {
  const [fileContent, setFileContent] = useState<string | null>(null);

  function uploadFile(event: React.ChangeEvent<HTMLInputElement>) {
    const fileInput = event.target;

    // Check if a file is selected
    if (fileInput.files && fileInput.files.length > 0) {
      const file = fileInput.files[0];
      const reader = new FileReader();

      reader.onload = function (event: any) {
        const content = event.target.result as string;
        setFileContent(content);
      };

      // Read the file as text
      reader.readAsText(file);
    } else {
      alert("No file selected.");
    }
  }

  function processFile() {
    if (fileContent) {
      const finalContent = {
        message: fileContent,
      }
      console.log(finalContent);
      // You can add your file processing logic here
    } else {
      alert("No file content to process.");
    }
  }

  return (
    <>
      <section className="flex flex-row justify-center">
        <div className="flex flex-col items-center">
          <h1 className="font-semibold text-gray-600 mt-2">or choose a file</h1>
          <input
            id="fileInput"
            type="file"
            accept=".txt, .pdf"
            onChange={(event) => uploadFile(event)}
            placeholder="Enter your text here"
            className="border p-2 rounded-md mt-5 transition-all duration-300 ease-in-out"
          />
          <div className="mt-4">
            <button onClick={() => processFile()} className="bg-indigo-500 text-white py-2 px-4 rounded-md mb-10">
              Process File

            </button>
          </div>
        </div>
      </section>
      <hr className="mr-20 ml-20 md:mr-48 md:ml-48 h-0.5" />
    </>
  );
}

