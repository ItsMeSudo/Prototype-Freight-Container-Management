import { Config, Links } from "../sidebar/config";
import { Button } from "../ui/button";
import { Link } from "react-router-dom";
export default function OnMobile(): JSX.Element {
  return (
    <>
      <div className="min-h-screen w-[100%] bg-[#292929] space-y-6 rounded-r-md flex flex-col">
        <div className="bg-[#202020] w-[100%] h-[125px] rounded-b-md flex flex-col justify-center items-center">
          <p className="font-inter-light">PFC UI</p>
        </div>
        <div>
          <div className="flex flex-col items-center justify-center space-y-4">
            {Links.map((item: Config) => (
              <>
                <Link to={item.RedirectTo}>
                  <Button className="bg-[#202020] hover:bg-[#202020]/50 w-[80%] text-white flex flex-row justify-start items-center">
                    <div className="flex flex-col w-[100%] justify-center items-center">
                      <item.Icon className=" h-5 w-5" />
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
