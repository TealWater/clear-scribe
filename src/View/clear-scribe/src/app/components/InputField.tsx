"use client"

import { useState } from "react";

export default function InputField() {
  const [descriptionValue, setDescriptionValue] = useState("");

  const handleInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    setDescriptionValue(event.target.value);
  };

  const handleButtonClick = () => {
    if (descriptionValue.trim() === "") {
      console.log("error");
      return;
    }
    const inputValue = {
      message: descriptionValue,
    };

    fetch("http://localhost:8080/send", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(inputValue),
    })
      .then((response) => response.json())
      .then((data) => {

        console.log(data);
      })
      .catch((error) => {
        console.error("Error:", error);
      });

    setDescriptionValue("");
  };

  return (
    <>
      <section className="flex flex-row justify-center">
        <div className="flex flex-col items-center">
          <textarea
            value={descriptionValue}
            onChange={handleInputChange}
            placeholder="Enter your text here"
            className="border p-2 rounded-md w-64 md:w-80 lg:w-96 h-48">
            </textarea>
          <div className="mt-4">
            <button onClick={handleButtonClick} className=" bg-cyan-500 text-white py-2 px-4 rounded-md mb-10">
              Process
            </button>
          </div>
        </div>
      </section>
      <hr />
    </>
  );
}
