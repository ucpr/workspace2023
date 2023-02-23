test1 = 100;
test2 = 200;
test3 = test1 + test2;
test4 = "AIUEO";
test5 = { 1, 2, 3, 4, 5 };
test6 = { a = 1, b = 2 };

-- print("Hello World")

function fibonacci(val)
  if val <= 0 then
    return 0;
  elseif val == 1 then
    return 1;
  end
  return fibonacci(val - 1) + fibonacci(val - 2)
end
