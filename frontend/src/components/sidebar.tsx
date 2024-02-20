import React from "react";
import OnDesktop from "./sidebar/OnDesktop";
import OnMobile from "./sidebar/OnMobile";

export default function Sidebar() : JSX.Element {
    const [width, updateWidth] = React.useState<number>(0);

    function OnUpdate() {
        if (width >= 0) {
            updateWidth(window.innerWidth)
        }
    }

    React.useEffect(() => {
        window.addEventListener("resize", () => {
            updateWidth(window.innerWidth)
        })

        OnUpdate()

    }, [])

    return (
        <>
            {width > 768 ? <>
                <OnDesktop/>
            </> : <>
                <OnMobile/>
            </>}
        </> 
    )
}