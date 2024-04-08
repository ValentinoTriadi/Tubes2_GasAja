"use client";

import { useState } from "react";
import { InputForm } from "./input-form";
import { Card, CardContent, CardHeader } from "./ui/card";
import { DisplayCard, ResultProps } from "./display-card";



export interface DataProps {
    Algorithm : string;
    StartKeyword: string;
    SearchKeyword: string;
    MaxIteration: number | undefined;
    Language: string;
}


export const MainCard = () => {

    const [data, setData] = useState<DataProps>({
        Algorithm: "",
        StartKeyword: "",
        SearchKeyword: "",
        MaxIteration: undefined,
        Language: ""
    });
    const [result, setResult] = useState<ResultProps | null>(null);

    return (
        <Card className="w-full sm:w-[99%] md:w-[97%] lg:w-[95%] bg-card rounded-md bg-clip-padding backdrop-filter backdrop-blur-md bg-opacity-70 border border-border z-0">
            <CardHeader>
                <h1 className="text-3xl font-bold">Welcome to <strong>GasAja</strong> WikiRace!</h1>
                <hr className="border-t-2 border-primary mt-2 w-full" />
            </CardHeader>
            <CardContent>
                <div className="flex gap-5 justify-start items-start w-full">
                    <InputForm setData={setData} setResult={setResult}></InputForm>
                    <DisplayCard data={data} result={result} ></DisplayCard>
                </div>
            </CardContent>
        </Card>
    );
};