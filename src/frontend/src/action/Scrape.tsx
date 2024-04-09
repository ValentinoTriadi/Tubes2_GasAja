"use server";

import { DataProps } from "@/components/MainCard";
import axios from "axios";
import { env } from "process";

export const scrape = async (data: DataProps) => {

    const body = {
        keyword: data.SearchKeyword,
        start: data.StartKeyword,
        limit: data.MaxIteration,
        lang: data.Language
    }
    const response = await axios.post(
        process.env.BACKEND_BASE + "/api/scrape",
        body,
        {
            headers: {
                'Content-Type': 'application/json'
            }
        }
    );
    console.log(response.data);
    return response.data;
}