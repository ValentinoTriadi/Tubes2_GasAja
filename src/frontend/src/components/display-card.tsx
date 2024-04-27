import Link from "next/link";
import { DataProps } from "./MainCard";
import { Card, CardContent, CardHeader } from "./ui/card";
import { ScrollArea } from "./ui/scroll-area";

interface ResultEntity {
    title : string;
    url : string;
}

export interface ResultProps {
    Webs : ResultEntity[][];
    Time : number;
    Total: number;
}

interface props {
    data : DataProps;
    result : ResultProps | null;
}

export const DisplayCard = ({ data, result } : props) => {
    const {Algorithm, StartKeyword, SearchKeyword, Language} = data;
    return (
        <div className="flex grow flex-wrap gap-5 max-w-full w-2/3">
            <Card className="min-w-[500px] max-w-full w-1/2 grow bg-popover rounded-md bg-clip-padding backdrop-filter backdrop-blur-md ">
                <CardHeader className="text-3xl font-bold">
                    Result
                </CardHeader>
                <CardContent>
                    <hr className="border-t-2 border-[var(--blue-11)] my-2 w-full " />
                    <div>
                        <h1><strong className="text-muted-foreground">Algorithm:</strong> {Algorithm === "ids" ? "Iterative Deepening Search" : Algorithm === "bfs" ? "Breadth First Search" : ""}</h1>
                        <h1><strong className="text-muted-foreground">Start Keyword:</strong> {StartKeyword}</h1>
                        <h1><strong className="text-muted-foreground">Search Keyword:</strong> {SearchKeyword}</h1>
                    </div>
                    <hr className="border-t-2 border-[var(--blue-11)] my-2 w-full" />
                    <ScrollArea className="max-h-full h-[400px]">
                        {result && result.Webs.map((webs, index) => {
                            return (
                                <div key={index} className="flex flex-col">
                                    <h1 className="text-xl font-bold my-2">Result {index + 1}</h1>
                                    <div className="flex flex-col">
                                        {webs.map((web, index) => {
                                            return (
                                                <div key={index} className="flex flex-wrap">
                                                    <h1 className="w-fit mr-1">{"->"} {web.title} </h1>
                                                    <Link href={ "https://" + Language + ".wikipedia.org" + web.url} className="text-ctextbase hover:text-ctexthover"> (https://{Language}.wikipedia.org{web.url})</Link>
                                                </div>
                                            );
                                        })}
                                    </div>
                                </div>
                            );
                        })}
                    </ScrollArea>
                </CardContent>
            </Card>
            <div className="flex grow flex-wrap gap-5 max-w-full w-1/6 min-w-[250px] h-fit">
                <Card className="w-fit h-fit grow bg-popover rounded-md bg-clip-padding backdrop-filter backdrop-blur-md py-5">
                    <CardContent className="px-4 py-2">
                        <div className="flex items-center ">
                            <h1><strong className="text-[var(--blue-11)]">Time:</strong> {result ? result.Time : 0}</h1>
                        </div>
                    </CardContent>
                </Card>
                <Card className="w-fit h-fit grow bg-popover rounded-md bg-clip-padding backdrop-filter backdrop-blur-md py-5">
                    <CardContent className="px-4 py-2">
                        <div className="flex items-center ">
                            <h1><strong className="text-[var(--blue-11)]">Total artikel yang dilalui:</strong> {result ? result.Webs[0].length -1 : 0}</h1>
                        </div>
                    </CardContent>
                </Card>
                <Card className="w-fit h-fit grow bg-popover rounded-md bg-clip-padding backdrop-filter backdrop-blur-md py-5">
                    <CardContent className="px-4 py-2">
                        <div className="flex items-center ">
                            <h1><strong className="text-[var(--blue-11)]">Total artikel yang dicari:</strong> {result ? result.Total : 0}</h1>
                        </div>
                    </CardContent>
                </Card>
            </div>
        </div>
    );
}