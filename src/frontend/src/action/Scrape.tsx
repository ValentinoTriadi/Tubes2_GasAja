"use server";

import { DataProps } from "@/components/MainCard";
import axios from "axios";

export const scrape = async (data: DataProps) => {

    const body = {
        keyword: data.SearchKeyword,
        start: data.StartKeyword,
        limit: data.MaxIteration
    }
    const response = await axios.post(
        "http://localhost:8000/api/scrape",
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