import Sidebar from "@/components/sidebar";
import Logo from "../../public/assets/logo.png";

export default function Dashboard() : JSX.Element {
    return (
        <>
            <div className="flex flex-row space-x-1">
                <div className="md:w-[10%] w-[25%]">
                    <Sidebar/>
                </div>
                <div className="flex md:w-[90%] w-[75%] mx-auto flex-col justify-center items-center">
                    <img src={Logo} className="rounded-full w-32 h-32"/>
                    <p className="font-inter-light break-normal text-[20px] text-center">Welcome to the Prototype Freight Container UI </p>
                </div>
            </div>
        </>
    )
}