import ADD_CSV from "@/components/Import/CSV";
import ADD_JSON from "@/components/Import/JSON";
import Sidebar from "@/components/sidebar";

export default function Import() {
  return (
    <>
      <div className="flex flex-row space-x-1">
        <div className="md:w-[10%] w-[25%]">
          <Sidebar />
        </div>
        <div className="flex md:w-[90%] w-[75%] mx-auto flex-col justify-center items-center">
          <div className="flex flex-col md:flex-row space-x-0 space-y-4 md:space-y-0 md:space-x-4">
            <div>
              <ADD_CSV />
            </div>
            <div>
              <ADD_JSON />
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
