export default function Footer(){

    return (
        <>
        <footer className="bg-slate-50 text-black py-10 mt-10 bottom-0 w-full">
        <div className="container mx-auto px-4 lg:px-8">
        <div className="grid grid-cols-1 lg:grid-cols-4 gap-8">
          <div className="flex justify-start "><svg className=" w-36 h-56" xmlns="http://www.w3.org/2000/svg" height="16" width="16" viewBox="0 0 512 512"><path d="M160 96a96 96 0 1 1 192 0A96 96 0 1 1 160 96zm80 152V512l-48.4-24.2c-20.9-10.4-43.5-17-66.8-19.3l-96-9.6C12.5 457.2 0 443.5 0 427V224c0-17.7 14.3-32 32-32H62.3c63.6 0 125.6 19.6 177.7 56zm32 264V248c52.1-36.4 114.1-56 177.7-56H480c17.7 0 32 14.3 32 32V427c0 16.4-12.5 30.2-28.8 31.8l-96 9.6c-23.2 2.3-45.9 8.9-66.8 19.3L272 512z"/></svg></div>
          <nav className="text-center  lg:text-left">
            <ul className="space-y-4">
              <h1 className="font-bold text-blue-500 brightness-90 underline decoration-blue-400">Site</h1>
              <li> <a href="/" className="hover:underline">Home</a></li>
              <li><a href="/history" className="hover:underline" >Your Files</a></li>
              <li><a href="/about" className="hover:underline">What do we do?</a></li>
            </ul>
          </nav>
          <div className="text-center lg:text-left">
            <h1 className="text-blue-500 brightness-90 underline decoration-blue-400 font-bold mb-4">Get in touch</h1>
            <h2 className="leading-loose">
              Discord: <br />
              Instagram: <br />
              Email: <br />
            </h2>
          </div>
          <div className="text-center lg:text-left">
            <h1 className="text-blue-500 brightness-90 underline decoration-blue-400 font-bold mb-4">Tools</h1>
            <h2 className="leading-loose font-semibold"> 
              Text processor: <br /> <p className="font-light text-sm ">A tool for converting text with complex words into a more simpler way.</p>
              File processor: <br /> <p className="font-light text-sm">Upload a file, we parse it then retreive the complex words and replace them for you.</p>
            </h2>
          </div>
        </div>
      </div>
      <div className="container mx-auto px-4 lg:px-8 mt-14 text-center lg:text-left">
        <p className="text-sm text-gray-500">&copy; 2023 ClearScribe. All rights reserved.</p>
      </div>
    </footer>   
        </>
    );
}