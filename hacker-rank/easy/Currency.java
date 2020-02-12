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

import java.text.NumberFormat;
import java.util.LinkedHashMap;
import java.util.Locale;
import java.util.Map;
import java.util.Scanner;

/**
 * Class that is executed in hacker rank website as solution.
 *
 * @author  Iván Camilo Sanabria (icsanabriar@googlemail.com)
 * @since   1.0.0
 */
public class Currency {

    /**
     * Main function provided by hacker rank website.
     *
     * @param args Arguments of the program.
     */
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);

        double payment = scanner.nextDouble();
        scanner.close();

        Map<String, Locale> locates = new LinkedHashMap<>();

        // Build proper instances of Locale for given countries.
        locates.put("US", Locale.US);
        locates.put("India", new Locale("en", "IN"));
        locates.put("China", Locale.CHINA);
        locates.put("France", Locale.FRANCE);

        // Iterate over each Locale and print the result.
        locates.forEach((k, v) -> {

            NumberFormat nf = NumberFormat.getCurrencyInstance(v);
            nf.setMaximumFractionDigits(2);

            String result = nf.format(payment);

            System.out.println((k + ": " + result));
        });
    }
}