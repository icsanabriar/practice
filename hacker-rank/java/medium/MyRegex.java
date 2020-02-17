/*
 * Copyright (C) 2020 Iván Camilo Sanabria.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import java.util.Scanner;

/**
 * Class that is executed in hacker rank website as solution.
 *
 * @author  Iván Camilo Sanabria (icsanabriar@googlemail.com)
 * @since   1.0.0
 */
public class MyRegex {

    /**
     * Pattern to identify numbers between 0 to 255, allowing 0 on the left side. (001) (011)
     */
    private static final String ACCEPTANCE_IP = "([01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])";

    /**
     * Pattern to identify IPs like 011.123.123.123.
     */
    private static final String PATTERN = ACCEPTANCE_IP + "\\." + ACCEPTANCE_IP + "\\." + ACCEPTANCE_IP + "\\." + ACCEPTANCE_IP;

    /**
     * Main function provided by hacker rank website.
     *
     * @param args Arguments of the program.
     */
    public static void main(String[] args) {

        Scanner in = new Scanner(System.in);

        while (in.hasNext()) {
            String IP = in.next();
            System.out.println(IP.matches(MyRegex.PATTERN));
        }
    }
}
