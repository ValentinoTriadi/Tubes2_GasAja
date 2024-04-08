"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import axios from "axios"
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


const numberSchema = z.preprocess((val) => {
	if (typeof val === 'string') return parseInt(val, 10);
	return val;
}, z.number().int().min(1, { message: "Minimum iteration is 1"}).max(100, { message: "Maximum iteration is 100" }))

const formSchema = z.object({
  StartKeyword: z.string().min(1, { message: "Start keyword must be filled"}).max(50, { message: "Start keyword maximum is 50 characters" }),
  SearchKeyword: z.string().min(1, { message: "Search keyword must be filled"}).max(50, { message: "Search keyword maximum is 50 characters" }),
  // MaxIteration: z.number().int().min(1, { message: "Minimum iterasi adalah 1"}).max(100, { message: "Maximum iterasi adalah 100" }),
  MaxIteration: numberSchema,
  Language: z.string().min(1, { message: "Language must be filled" }).max(2, { message: "Language maximum is 2 characters" })
})

interface props {
  setData: React.Dispatch<React.SetStateAction<any>>;
  setResult: React.Dispatch<React.SetStateAction<any>>;
}


export function InputForm({setData, setResult} : props) {

  const [algorithm, setAlgorithm] = useState<string>("bfs");

  const [isPending, startTransition] = useTransition()
  const [isLoading, setIsLoading] = useState(false)

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      StartKeyword: "",
      SearchKeyword: "",
      MaxIteration: 1,
      Language: ""
    },
    mode: "all"
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    const data : DataProps = {
      Algorithm: algorithm,
      StartKeyword: values.StartKeyword,
      SearchKeyword: values.SearchKeyword,
      MaxIteration: algorithm === "ids" ? values.MaxIteration : 1,
      Language: values.Language,
    }
    setIsLoading(true);
    setData(data);
    startTransition(() => {
      scrape(data).then((res) => {
        setResult(res);
      })
			setIsLoading(false)
		})
  }

  // 3. Render your form.
  return (
    <div className="w-1/3">
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
          {algorithm === "ids" && <FormField
            control={form.control}
            name="MaxIteration"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Maximum Iteration</FormLabel>
                <FormControl>
                <Input placeholder="1" type="number" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />}
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
          <Button type="submit" disabled={isLoading}>Start</Button>
        </form>
      </Form>
    </div>

  )
}