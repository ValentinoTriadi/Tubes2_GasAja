"use server";

import { DataProps } from "@/components/MainCard";
import axios from "axios";

export const scrape = async (data: DataProps) => {

    const body = {
        keyword: data.SearchKeyword,
        start: data.StartKeyword,
        lang: data.Language
    }

    const fullURL = process.env.BACKEND_BASE + "/api/scrape/" + data.Algorithm;
    // console.log(fullURL);
    // console.log(body);
    const response = await axios.post(
        fullURL,
        body,
        {
            headers: {
                'Content-Type': 'application/json'
            }
        }
    );
    // console.log(response.data);
    // console.log(response.data.Webs)
    // // console log response.data.webs
    // for (let i = 0; i < response.data.Webs.length; i++) {
    //     for (let j = 0; j < response.data.Webs[i].length; j++) {
    //         console.log(response.data.Webs[i][j])
    //     }
    // }
    return response.data;
}