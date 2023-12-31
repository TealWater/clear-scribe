"use client"

import { useState } from "react";
let url =  process.env.NEXT_PUBLIC_BACKEND_API;

export default function InputField() {
  const [descriptionValue, setDescriptionValue] = useState("");

  const handleInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {

    // get the value
    setDescriptionValue(event.target.value);
  };

  const handleButtonClick = () => {

    // handles spaces to see if something was sent
    if (descriptionValue.trim() === "") {
      alert("Error! There is nothing to process")
      return;
    }
    // make it an object
    const inputValue = {
      message: descriptionValue,
    };


    // send to backend via POST and convert it to json
    fetch(`${url}/send`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(inputValue),
    })
      .then((response) => response.json())
      .then((data) => {
        if(data){
          const button = document.getElementById("sucessMessage");
          button?.classList.remove("hidden");
          const button2 = document.getElementById("errorMessage");
          button2?.classList.add("hidden");
        } else {
          const button = document.getElementById("errorMessage");
          button?.classList.remove("hidden");
          const button2 = document.getElementById("sucessMessage");
          button2?.classList.add("hidden");
        }
        console.log(data);
      })
      .catch((error) => {
        console.error("Error:", error);
      });


      // clear textarea value
    setDescriptionValue("");
  };

  return (
    <>
      <section className="flex flex-row justify-center">
        <div className="flex flex-col items-center">

          <h1 className="text-xl font-bold mb-4">Enter your text:</h1>
          <textarea
            name="inputArea"
            value={descriptionValue}
            onChange={handleInputChange}
            placeholder="Enter your text here"
            className="border p-2 rounded-md w-64 md:w-80 lg:w-96 h-48">
            </textarea>
          <div className="mt-4">

            <button onClick={handleButtonClick} className=" bg-indigo-500 text-white py-2 px-4 rounded-md mb-5 duration-200 hover:scale-105">
              Process Text
            </button>
          </div>
          <p className='text-green-500 mb-5 hidden' id='sucessMessage'>Text processed. Please click on<a className='underline' href='/history'> 'Files'</a></p>
          <p className='text-red-500 mb-5 hidden' id='errorMessage'>File couldn't be processed, try again later</p>
        </div>
      </section>

    </>
  );
}
