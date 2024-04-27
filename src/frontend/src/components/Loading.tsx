export const Loading = () => {
    return (
        <div className="w-full h-full fixed top-0 left-0 flex flex-col items-center justify-center bg-popover z-10">
            <svg width="100" height="100" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg" className="animate-spin">
                <circle cx="50" cy="50" r="20" fill="none" strokeWidth="2" stroke="var(--foreground)" strokeDasharray="31.42 31.42" strokeLinecap="round" transform="rotate(0 50 50)"></circle>
            </svg>
            <p className="text-xl animate-bounce" >{"Loading..."}</p>
        </div>
    )
}