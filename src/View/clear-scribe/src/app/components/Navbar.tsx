import Image from "next/image";

export default function Navbar() {

  return (
    <nav className="flex justify-between items-center p-5 ">
      <div className="text-black font-bold">Logo</div>
      <div>
        <ul className="flex space-x-3 text-sm md:text-base md:space-x-5 mr-10 md:mr-20 lg:mr-40  text-black">
          <li>
            <a href="/" className=" hover:text-gray-500">Home</a>
          </li>
          <li>
            <a href="/history" className=" hover:text-gray-500">History</a>
          </li>
          <li>
            <a href="/about" className=" hover:text-gray-500">About</a>
          </li>
        </ul>
      </div>
    </nav>
  );
}
