import React from "react";

export default function MainText() {
  return (
    <div className="flex flex-col text-center items-center bg-gradient-to-b from-purple-500 to-indigo-800 text-white overflow-hidden p-8 md:p-16 mb-8">
      <h1 className="font-bold text-4xl lg:text-6xl">
        We make texts easier.
      </h1>
      <h1 className="font-bold text-5xl lg:text-7xl mt-2 text-blue-500">
        CLEAR SCRIBE
      </h1>
      <p className="font-light text-gray-300 text-sm lg:text-xl max-w-sm lg:max-w-2xl mt-4">
        Clear Scribe is your ultimate tool for converting complex text into clear, concise, and easily understandable language. Whether you&lsquo;re a student grappling with a challenging textbook, a professional deciphering intricate documents, or simply someone who values straightforward communication, Clear Scribe has got you covered.
      </p>
    </div>
  );
}
