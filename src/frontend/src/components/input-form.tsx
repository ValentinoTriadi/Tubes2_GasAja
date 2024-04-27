"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import React, { useState, useTransition } from "react"
import { 
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue 
} from "./ui/select"
import { DataProps } from "./MainCard"
import { scrape } from "@/action/Scrape"
import { Loading } from "./Loading"


const formSchema = z.object({
  StartKeyword: z.string().min(1, { message: "Start keyword must be filled"}).max(50, { message: "Start keyword maximum is 50 characters" }),
  SearchKeyword: z.string().min(1, { message: "Search keyword must be filled"}).max(50, { message: "Search keyword maximum is 50 characters" }),
  Language: z.string().min(1, { message: "Language must be filled" }).max(2, { message: "Language maximum is 2 characters" })
})

interface props {
  setData: React.Dispatch<React.SetStateAction<any>>;
  setResult: React.Dispatch<React.SetStateAction<any>>;
}


export function InputForm({setData, setResult} : props) {

  const [algorithm, setAlgorithm] = useState<string>("bfs");
  const [isPending, startTransition] = useTransition()

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      StartKeyword: "",
      SearchKeyword: "",
      Language: ""
    },
    mode: "all"
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    const data : DataProps = {
      Algorithm: algorithm,
      StartKeyword: values.StartKeyword,
      SearchKeyword: values.SearchKeyword,
      Language: values.Language,
    }
    // setIsLoading(true);
    setData(data);
    startTransition(() => {
      scrape(data).then((res) => {
        setResult(res);
      })
			// setIsLoading(false)
		})
  }

  // 3. Render your form.
  return (
    <div className="grow min-w-[32%] w-[32%]">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <div className="flex gap-5 items-center">
            <h2 className="text-xl font-semibold text-muted-foreground w-fit" >Algorithm:</h2>
            <Select onValueChange={(value) => setAlgorithm(value)} required>
              <SelectTrigger>
                <SelectValue placeholder="Select the Algorithm"/>
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value="bfs">Breadth First Search</SelectItem>
                  <SelectItem value="ids">Iterative Deepening Search</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>
          <FormField
            control={form.control}
            name="StartKeyword"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Start Keyword</FormLabel>
                <FormControl>
                <Input placeholder="Kucing" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
            />
          <FormField
            control={form.control}
            name="SearchKeyword"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Search Keyword</FormLabel>
                <FormControl>
                <Input placeholder="Bahasa Jawa" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="Language"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Language</FormLabel>
                <FormControl>
                <Select onValueChange={field.onChange} defaultValue={field.value} {...field} required>
                  <SelectTrigger>
                    <SelectValue placeholder="Select the Language"/>
                  </SelectTrigger>
                  <SelectContent>
                    <SelectGroup>
                      <SelectItem value="id">Bahasa Indonesia</SelectItem>
                      <SelectItem value="en">English</SelectItem>
                    </SelectGroup>
                  </SelectContent>
                </Select>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" disabled={isPending}>Start</Button>
        </form>
      </Form>
      {isPending && <Loading/>}
    </div>
  )
}