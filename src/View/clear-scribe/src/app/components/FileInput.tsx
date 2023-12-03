export default function FileInput() {

    return (
        <>
        <section className="flex flex-row justify-center">
        <div className="flex flex-col items-center">
          <input
           type="file"
            placeholder="Enter your text here"
            className="border p-2 rounded-md mt-5 ">
            </input>
          <div className="mt-4">
            <button className=" bg-cyan-500 text-white py-2 px-4 rounded-md mb-10">
              Process
            </button>
          </div>
        </div>
      </section>
        </>
    );
}