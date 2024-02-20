import { Config, Links } from "../sidebar/config";
import { Button } from "../ui/button";
import { Link } from "react-router-dom";
export default function OnDesktop(): JSX.Element {
  return (
    <>
      <div className="min-h-screen w-[100%] bg-[#292929] space-y-6 rounded-r-md flex flex-col">
        <div className="bg-[#202020] w-[100%] h-[125px] rounded-b-md flex flex-col justify-center items-center">
          <p className="font-inter-light text-center">Prototype Freight Container</p>
        </div>
        <div>
          <div className="flex flex-col items-center space-y-4">
            {Links.map((item: Config) => (
              <>
                <Link to={item.RedirectTo}>
                  <Button className="bg-[#202020] space-x-1 hover:bg-[#202020]/50 min-w-[80%] max-w-[100%] text-white hover: flex flex-row justify-start items-center">
                    <div className="">
                      <item.Icon className=" h-5 w-5" />
                    </div>
                    <div className="flex flex-col justify-center items-center w-[90%]">
                      <p>{item.Title}</p>
                    </div>
                  </Button>
                </Link>
              </>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}
