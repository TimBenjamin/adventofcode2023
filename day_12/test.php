<?php
if (!function_exists('str_starts_with')) {
    function str_starts_with(string $haystack, string $needle): bool {
        $lh = strlen($haystack);
        $ln = strlen($needle);
        if ($ln > $lh)
            return false;
        $sh = substr($haystack, 0, $ln);
        return $sh === $needle;
    }
}

if (!function_exists('str_ends_with')) {
    function str_ends_with(string $haystack, string $needle): bool {
        $lh = strlen($haystack);
        $ln = strlen($needle);
        if ($ln > $lh)
            return false;
        $sh = substr($haystack, $lh - $ln, $ln);
        return $sh === $needle;
    }
}

/**
 * @param string $str
 * @param string $separator
 * @return int[]
 */
function getNumbers(string $str, string $separator = ' '): array
{
    $numbers = [];
    $parts = explode($separator, $str);
    foreach ($parts as $part)
    {
        $part = trim($part);
        if (strlen($part) > 0 && is_numeric($part))
            $numbers[] = intval($part);
    }
    return $numbers;
}


/**
 * @param int $day
 * @return resource
 */
function getInputFile(int $day) {
    $dayName = 'day' . str_pad("$day", 2, '0', STR_PAD_LEFT);
    // $pathParts = [__DIR__, '..', 'advent-of-code-2023-data', $dayName, 'input.txt'];
    $path = "./input.txt"; // join(DIRECTORY_SEPARATOR, $pathParts);
    $file = fopen($path, 'r');
    if ($file === false)
        throw new RuntimeException('cannot open file');
    return $file;
}

$file = getInputFile(12);

$sum = 0;
while (($line = fgets($file)) !== false)
{
    $line = trim($line);
    $parts = explode(' ', $line);
    if (count($parts) !== 2)
        continue;
    $springs = $parts[0];
    $springs = "$springs?$springs?$springs?$springs?$springs";
    $damaged = getNumbers($parts[1], ',');
    $damaged = array_merge($damaged, $damaged, $damaged, $damaged, $damaged);
    $cPossibles = getPossiblesCount($springs, null, ...$damaged);
    $sum += $cPossibles;
    print ("$cPossibles\t - $sum\n");
}

global $cache;
$cache = [];

function getPossiblesCount(string $springs, ?int $remaining = null, int ... $damaged): int
{
    global $cache;
    $key = $springs.'-'.join('-', $damaged);
    if (isset($cache[$key])) {
        return $cache[$key];
    }
    $d = array_shift($damaged);
    if ($d === null || $d < 1) {
        $index = stripos($springs, '#');
        if ($index !== false)
            return 0;
        return 1;
    }

    if ($remaining === null)
    {
        $remaining = 0;
        foreach ($damaged as $i)
            $remaining += $i;
    } else {
        $remaining -= $d;
    }

    $sSprings = strlen($springs);
    $possibles = 0;
    $iMax = $sSprings - $d - $remaining + 1;
    for ($i = 0; $i < $iMax; ++$i)
    {
        if ($springs[$i] === '#')
            $iMax = min($i + 1, $iMax);
        $ok = $i + $d === $sSprings || $springs[$i + $d] === '.' || $springs[$i + $d] === '?';
        $ok &= $i === 0 || $springs[$i - 1] === '.' || $springs[$i - 1] === '?';

        for ($j = $i; $j < $sSprings && $j < $i + $d && $ok; ++$j)
        {

            if ($springs[$j] !== '?' && $springs[$j] !== '#')
                $ok = false;
        }
        if ($ok)
        {
            if ($i + $d === $sSprings) {
                $possibles += 1;
            } else {
                $subSprings = substr($springs, $i + $d + 1);
                $subPossibles = getPossiblesCount($subSprings, $remaining, ... $damaged);
                $possibles += $subPossibles;
            }
        }
    }
    $cache[$key] = $possibles;
    return $possibles;
}