"use client";

import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const formSchema = z.object({
  email: z.string().email(),
  email_head: z.string(),
  email_body: z.string(),
});

export function EmailForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      email_body: "",
      email_head: "",
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
  }

  return (
    <Form {...form}>
      <form
        onSubmit={form.handleSubmit(onSubmit)}
        className="space-y-8 container"
      >
        <div className="flex flex-row py-3">
          <div className="w-3/4 px-5 py-2">
            <FormField
              control={form.control}
              name="email_head"
              render={({ field }) => (
                <FormItem className="pb-5">
                  <FormLabel>Email Subject</FormLabel>
                  <FormControl>
                    <Input placeholder="" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="email_body"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Body</FormLabel>
                  <FormControl>
                    <Textarea
                      className="min-h-[300px] md:min-h-[500px] lg:min-h-[500px]"
                      placeholder="Long Time No See"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            ></FormField>
          </div>
          <div className=" flex flex-col items-center w-1/4">
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem className="py-5 w-full lg:w-3/4">
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      type="email"
                      className=""
                      placeholder="abc@xyz.com"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage /> 
                </FormItem>
              )}
            />
            <Button className="w-full lg:w-3/4 pt-2" type="submit">
              Submit
            </Button>
          </div>
        </div>
      </form>
    </Form>
  );
}
